package middleware

import (
	"log/slog"
	"strings"
	
	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		// do not log swagger requests not to clutter the logs
		if c.Method() == "GET" && strings.HasPrefix(c.Path(), "/swagger/") {
			return err
		}
		slog.Info(
			"request",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"response", string(c.Response().Body()),
		)
		return err
	}
}
