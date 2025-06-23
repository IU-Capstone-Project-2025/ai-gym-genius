package authorization

import (
	"admin/internal/authorization/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/auth", handlers.LoginHandler)
}
