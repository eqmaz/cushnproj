package middleware

import (
	"time"

	"cushon/internal/application/cfg"
	//"cushon/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// FiberRateLimiter - very general, just limits the number of requests per IP address for Fiber router
// This would be the last line of defense; we'd expect the gateway to handle most of the rate limiting,
// but it gives us application-level control as well
func FiberRateLimiter(cfg cfg.RateLimit) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        cfg.MaxRequests,
		Expiration: time.Duration(cfg.Timeframe) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				// Note, we could use our global Reply shape here as well
				"error":  "Rate limit exceeded",
				"result": nil,
			})
		},
	})
}

// GinRateLimiter - limits the number of requests per IP address for Gin router
// Placeholder, in case we want to swap to another REST framework
// it can be easily implemented
