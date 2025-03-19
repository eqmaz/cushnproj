package application

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"cushon/internal/application/cfg"
	"cushon/internal/router"
	"cushon/pkg/config"
	"cushon/pkg/database"
	"cushon/pkg/health"

	c "cushon/pkg/console"
	"cushon/pkg/e"
	"cushon/pkg/logger"
	. "cushon/pkg/optional"
)

// App is our main application struct and contains the configuration, logger and router
type App struct {
	Config   *cfg.Config
	Database *database.MySQL
	Logger   *logger.Logger
	Router   router.Router
	Health   *health.Health
}

// NewApp initializes the App with the necessary dependencies
func NewApp(logger *logger.Logger) *App {
	return &App{
		Logger: logger,
	}
}

// SetConfigs loads the application configuration from the defaults, cfg.json file and environment variables
func (app *App) SetConfigs() error {
	// Set up the error catalogue
	e.SetCatalogue(errorMap)

	configFilePath, err := config.GetConfigFilePath(None[string](), config.DefaultPreferences())
	if err != nil {
		c.Error("Error getting config file path: %s\n", err.Error())
		return err
	}

	var appConfig cfg.Config
	err = appConfig.LoadConfigs(configFilePath)
	if err != nil {
		c.Outf("Error loading configs: %s\n", err.Error())
		return err
	}

	// Check if console logging is enabled
	if appConfig.Console.Enabled == false {
		c.Info("Console messages suspended")
		c.Suspend()
	}

	app.Config = &appConfig

	// Print out loaded config
	c.Infof("Config loaded from %s", configFilePath)
	return nil
}

// getDatabaseConfig - Internal function to get the database configuration
func (app *App) getDatabaseConfig() (*database.DbConfig, error) {
	appDbCfg := &app.Config.Database
	// Check that the User and Pass are not empty
	// They would typically be set via the environment
	if appDbCfg.User == "" || appDbCfg.Pass == "" {
		return nil, e.FromCode("eDupNs")
	}

	dbConfig := database.DbConfig{
		Driver:   "mysql",
		Host:     appDbCfg.Host,
		Port:     appDbCfg.Port,
		Name:     appDbCfg.Name,
		User:     appDbCfg.User,
		Password: appDbCfg.Pass,

		// TODO - hard coded for now
		MaxOpenCon:      10,
		MaxIdleCon:      5,
		ConnMaxLifetime: 30 * time.Minute,
	}

	return &dbConfig, nil
}

// SetDatabase initializes the database connection
func (app *App) SetDatabase() *App {
	// Ensure configs are loaded at this point
	if app.Config == nil {
		c.Warn("No config initialized. Ensure configs are set up before starting server.")
		os.Exit(1)
	}

	// Ensure Logger is initialized
	if app.Logger == nil {
		c.Warn("No logger initialized. Ensure logger is set up before starting server.")
		os.Exit(1)
	}

	log := app.Logger

	// Set up the database configuration
	dbConfig, err := app.getDatabaseConfig()
	if err != nil {
		log.Error(err.Error(), nil)
		os.Exit(1)
	}

	// Initialize the database (MySQL implementation)
	db := database.NewMySQL(*dbConfig, log)
	if err := db.Connect(); err != nil {
		log.Fatal("Unable to start database:", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping database:", err)
	}
	if _, err := db.Exec("SELECT 1"); err != nil {
		log.Fatal("Unable to execute SELECT 1 query:", err)
	}

	// Set the database connection in App instance
	app.Database = db

	c.Success("Database test complete.")
	return app
}

// SetRoutes initializes the Router, route handlers and middleware
func (app *App) SetRoutes() *App {
	app.Router = router.NewFiberRouter(app.Logger, app.Config, app.Database.SqlxDb)
	c.Info("Using Fiber router")

	app.Router.RegisterMiddleware()
	app.Router.RegisterHandlers() // uses the generated API handlers
	return app
}

// MonitorSignals - listens for OS signals and exits the program gracefully
func (app *App) MonitorSignals() *App {
	// Create a channel to receive OS signals
	sigs := make(chan os.Signal, 1)

	// Notify the channel on receiving SIGINT (Ctrl+C) or SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to notify the program to exit
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		c.Out("Received: ", sig)
		done <- true
	}()

	go func() {
		<-done
		c.Out("Stopping server...")
		//app.Router.Stop()
		os.Exit(0)
	}()

	//c.Out("Press Ctrl+C to exit")
	return app
}

// StartHealth - starts the health check server
func (app *App) StartHealth() *App {
	var hConfig = app.Config.Health

	healthServer := health.New(health.Options{
		HTTPEnabled: hConfig.Http.Enabled,
		GRPCEnabled: hConfig.Grpc.Enabled,
		HTTPPort:    strconv.FormatUint(uint64(hConfig.Http.Port), 10),
		GRPCPort:    strconv.FormatUint(uint64(hConfig.Grpc.Port), 10),
		CheckFunction: func(ctx context.Context) (health.HealthResult, error) {

			// TODO - we can implement a sensible health check logic

			// Check the app database health
			if err := app.Database.Ping(); err != nil {
				return health.HealthResult{
					Status:  "ERROR",
					Message: "Database connection error",
				}, err
			}

			return health.HealthResult{Status: "OK"}, nil
		},
	})

	app.Health = healthServer

	healthServer.Start()

	return app
}

// Serve - starts the main business logic for the application
func (app *App) Serve() {
	// Ensure configs are loaded at this point
	if app.Config == nil {
		c.Warn("No config initialized. Ensure configs are set up before starting server.")
		return
	}

	// Ensure the Router is initialized
	if app.Router == nil {
		c.Warn("No router initialized. Ensure routes are set up before starting server.")
		return
	}

	// Get the REST API port from configs
	portStr := strconv.FormatUint(app.Config.Service.Port, 10)

	address := ":" + portStr // e.g. :8080
	err := app.Router.Serve(address)
	if err != nil {
		c.Warnf("Server could not listen on the given port")
		e.FromError(err).Print(0, 0)
		return
	}

	c.Successf("Application started on port " + portStr)
}

func GetErrorMap() map[string]string {
	return errorMap
}
