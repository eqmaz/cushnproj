package router

import (
	"cushon/internal/application/cfg"
	"cushon/internal/router/api"
	"cushon/internal/router/fiber"
	"cushon/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// Router interface defines the methods that a reasonable REST framework should implement
// Do not confuse that with "handlers", the actual functions that handle the requests.
// The idea is to make the router implementation interchangeable
// Today we've gone with Fiber, but we could swap, or even add Gin, or another framework side by side
type Router interface {
	RegisterMiddleware()
	RegisterHandlers()
	Serve(addr string) error
}

// NewFiberRouter creates a new instance of the FiberRouter
func NewFiberRouter(logger *logger.Logger, config *cfg.Config, db *sqlx.DB) Router {
	// Set up the Fiber configuration
	fiberConfig := fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},

		// Whether to not show the "Fiber" greeting message
		DisableStartupMessage: config.Service.DisableRouterGreeting,
	}

	// Create a new Fiber app with a custom error handler
	fiberApp := fiber.New(fiberConfig)
	restServer := api.NewRestServer(db, config.Service)

	return &fiberHandlers.FiberRouter{
		App:    fiberApp,
		Config: config,
		Logger: logger,
		Server: restServer,
	}
}

// Placeholder - if you wanted a choice of REST frameworks, you could add a Gin router here, and so on
// eg: NewGinRouter creates a new instance of the GinRouter
