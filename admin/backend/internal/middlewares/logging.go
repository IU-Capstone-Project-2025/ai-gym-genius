package middleware

import (
	"log/slog"
	"strings"
	
	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		response := ""
		path := c.Path()
		// do not log swagger or welcome page responses not to clutter the logs
		if !(c.Method() == "GET" && (strings.HasPrefix(path, "/swagger/") || path == "/")) {
			response = string(c.Response().Body())
		}
		slog.Info(
			"request",
			"method", c.Method(),
			"path", path,
			"status", c.Response().StatusCode(),
			"response", response,
		)
		return err
	}
}
