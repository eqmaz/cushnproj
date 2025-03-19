package main

import (
	"cushon/cmd/version"
	. "cushon/internal/application"
	c "cushon/pkg/console"
	"cushon/pkg/database"
	"cushon/pkg/e"
	"cushon/pkg/logger"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// This will remove verbose full-path log lines when printing backtraces and errors
	// We only need to show filenames relative to the project root
	e.SetFilePathTrimPoint("/cushon")

	version.PrintVersion()

	// Spawn a new logger instance
	appLogger := logger.NewLogger()

	// Create and boot the application
	app := NewApp(appLogger)

	// Load the application configuration
	// It will first look for CLI arguments passed in (--config=/path/to/config.json),
	// then it will look for a 'config.json' in the current working dir and executable dir.
	// If no file is found, the default values will be used.
	if err := app.SetConfigs(); err != nil {
		appLogger.Error(err.Error(), nil)
		c.Error("Error setting configs: %s", err.Error())
		return
	}

	// Graceful shutdown of database connection when application exits
	defer func(db *database.MySQL) {
		if err := db.Close(); err != nil {
			log.Errorf("Error closing database connection: %v", err) // Should never happen
		}
	}(app.Database)

	app.MonitorSignals(). // Monitor for SIGINT and SIGTERM signals to gracefully shutdown the server
				SetDatabase(). // Set up the database connection
				StartHealth(). // Start the health check server (if enabled)
				SetRoutes().   // Set up the routes for the application
				Serve()        // Start HTTP server and business logic
}
