// Exclude from normal build and unit tests
//go:build integration

package integration

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"cushon/internal/application/apperr"
	"cushon/internal/application/cfg"
	"cushon/internal/router"
	c "cushon/pkg/console"
	"cushon/pkg/database"
	"cushon/pkg/e"
	"cushon/pkg/logger"
)

// Super simple test runner
// In a real world app we'd have a TestApp struct that would
// do all the setup stuff, and we'd define "suites" of tests
// each testing a different domain

var BaseURL = "http://127.0.0.1:8080"
var testDB *database.MySQL
var testRouter router.Router

// getDatabaseConfig - Internal function to get the database configuration
func getDatabaseConfig(config *cfg.Config) (*database.DbConfig, error) {
	appDbCfg := config.Database
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

func runTestMigrations(db *database.MySQL) {
	// We haven't really got any yet
}

func tearDownTestDatabase(db *database.MySQL) {
	// in case we want to do some cleanup
}

// TestMain is the entry point for integration tests.
func TestMain(m *testing.M) {
	// This loads our app's error code catalog
	e.SetCatalogue(apperr.GetErrorMap())

	// Load config file (we could have a different one for testing)
	var config cfg.Config
	err := config.LoadConfigs("../../config/config.json")
	if err != nil {
		c.Outf("Error loading configs: %s\n", err.Error())
		return
	}

	// Setup logger instance
	testLogger := logger.NewLogger()
	//testLogger.Suspend()

	// Initialize test database
	dbConfig, _ := getDatabaseConfig(&config)
	testDB = database.NewMySQL(*dbConfig, testLogger)
	err = testDB.Connect()
	if err != nil {
		c.Error("Error connecting to database: " + err.Error())
		return
	}

	// Create the Fiber Router with test dependencies
	testRouter = router.NewFiberRouter(testLogger, &config, testDB.SqlxDb)
	testRouter.RegisterMiddleware()
	testRouter.RegisterHandlers()

	// Run migrations (not needed yet)
	// runTestMigrations(testDB)

	// Start the test server in a goroutine
	go func() {
		portStr := ":" + strconv.FormatUint(config.Service.Port, 10)
		c.Outf("Listening on port %s\n", portStr)
		if err := testRouter.Serve(portStr); err != nil {
			log.Fatalf("Failed to start test server: %v", err)
		}
	}()

	time.Sleep(1 * time.Second) // Wait for server to spin up

	// Run tests
	code := m.Run()

	// Cleanup test DB and shutdown
	tearDownTestDatabase(testDB)

	os.Exit(code)
}
