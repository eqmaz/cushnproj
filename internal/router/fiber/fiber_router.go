package fiberHandlers

import (
	"cushon/internal/application/cfg"
	"cushon/internal/router/api"
	"cushon/internal/router/middleware"
	c "cushon/pkg/console"
	l "cushon/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	App    *fiber.App
	Config *cfg.Config
	Logger *l.Logger
	Server api.ServerInterface
}

func (r *FiberRouter) RegisterMiddleware() {
	r.App.Use(middleware.FiberLogger(r.Logger))
	r.App.Use(middleware.FiberRateLimiter(r.Config.RateLimit))

	// Set Content-Type to JSON
	r.App.Use(func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Next()
	})

	// Ensure received body is JSON
	r.App.Use(func(c *fiber.Ctx) error {

		// check if it's a POST request
		if c.Method() == "POST" {
			if !c.Is("json") {
				return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
					"error": "Content-Type must be application/json",
				})
			}
		}

		return c.Next()
	})
}

func (r *FiberRouter) RegisterHandlers() {
	// Handle favicon requests from browsers explicitly
	r.App.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
	api.RegisterHandlers(r.App, r.Server)
}

func (r *FiberRouter) Serve(addr string) error {

	// Check if TLS is enabled
	tlsConfig := &r.Config.Service.Tls
	if tlsConfig.Enabled {
		c.Info("Starting REST service with TLS on port " + addr)
		return r.App.ListenTLS(addr, tlsConfig.Cert, tlsConfig.Key)
		//go func() {
		//    if err := r.App.ListenTLS(addr, tlsConfig.Cert, tlsConfig.Key); err != nil {
		//        r.Logger.Errorf("Failed to start HTTPS server: %v", err)
		//    }
		//}()
	}

	c.Info("Starting REST service [without TLS] on port " + addr)
	return r.App.Listen(addr)
	//go func() {
	//    if err := r.App.Listen(addr); err != nil {
	//        r.Logger.Errorf("Failed to start HTTPS server: %v", err)
	//    }
	//}()
}
