package api

import (
	"strconv"

	"cushon/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func getCtxLogger(ctx *fiber.Ctx) *logger.Logger {
	return ctx.Locals("logger").(*logger.Logger)
}

// getUserId - placeholder function to simulate getting the user ID from the request context
// In the real world we'd get it from a JWT or similar auth mechanism
func getUserId(ctx *fiber.Ctx) (uint64, error) {
	userIdStr := ctx.Get("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
