package health

import (
	"context"
	"net/http"

	c "cushon/pkg/console"
	"github.com/gofiber/fiber/v2"
)

// HTTPServer wraps a Fiber app.
type HTTPServer struct {
	app    *fiber.App
	health *Health
}

// NewHTTPServer creates a new HTTP server instance.
func NewHTTPServer(h *Health) *HTTPServer {
	fiberConfig := fiber.Config{
		DisableStartupMessage: true,
	}
	app := fiber.New(fiberConfig)
	srv := &HTTPServer{app: app, health: h}
	app.Get("/ping", srv.pingHandler)
	app.Get("/health", srv.healthHandler)
	return srv
}

// pingHandler - simple heartbeat response, just says "I'm alive" but doesn't check any services.
func (s *HTTPServer) pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong") // Probably a 200 OK would be sufficient, it's just for show
}

// healthHandler - runs the full health check and returns the result.
func (s *HTTPServer) healthHandler(c *fiber.Ctx) error {
	result, err := s.health.HealthCheck(context.Background())
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(result)
	}
	return c.JSON(result)
}

// Start runs the Healthcheck Fiber app.
func (s *HTTPServer) Start(port string) {
	c.Infof("HTTP health server listening on port %s", port)
	if err := s.app.Listen(":" + port); err != nil {
		c.Error(err.Error())
	}
}

// Stop shuts down the Fiber app.
func (s *HTTPServer) Stop() {
	err := s.app.Shutdown()
	if err != nil {
		c.Error(err.Error()) // Should never happen
		return
	}
}
