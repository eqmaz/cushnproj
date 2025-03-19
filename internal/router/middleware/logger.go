package middleware

import (
	"fmt"

	"cushon/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// makeContextLogger helper used by all routers' middleware to create a new logger with the path and IP
func makeContextLogger(baseLogger *logger.Logger, path, ip string, cid string) *logger.Logger {
	fields := map[string]interface{}{
		"path": path,
		"ip":   ip,
	}

	// Add the correlation ID if it exists
	if cid != "" {
		fields["cid"] = cid
	}
	return baseLogger.WithContextFields(fields)
}

// FiberLogger attaches a new contextual logger to the context, for Fiber router
func FiberLogger(baseLogger *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		// Try to capture a correlation ID from the request headers
		correlationID := c.Get("X-Correlation-Id") // capture header
		if correlationID != "" {
			c.Locals("cid", correlationID)
		}

		// Create a new logger with the path and IP
		ctxLogger := makeContextLogger(baseLogger, path, c.IP(), correlationID)

		// Don't log calls to favicon
		if path != "/favicon.ico" {
			ctxLogger.Info(fmt.Sprintf("req: %s %s", c.Method(), c.OriginalURL()), nil)
		}

		// Set the logger in the context so it's easily available to the handlers
		// Assumes the logger lifetime will outlive the context
		c.Locals("logger", ctxLogger)

		return c.Next()
	}
}

// GinLogger attaches a new contextual logger to the context, for Gin router
// Not implemented, but in case we want to switch out the router implementation
