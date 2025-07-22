package middleware

import (
	"admin/config"
	"admin/internal/models"

	"github.com/gofiber/fiber/v2"
)

// to be used after JWTMiddleware
func AdminOnly(c *fiber.Ctx) error {
	if config.C.AppEnv == "DEV" {
		return c.Next() // skip admin-only checks for development
	}
	role, ok := c.Locals(RoleKey).(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "unauthorized or malformed role claim",
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "admin-only route",
		})
	}

	return c.Next()
}
