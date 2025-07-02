package exercises

import (
	"admin/internal/exercises/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Post("/exercises/", handlers.AddExercise)
	app.Delete("/exercises/:id", handlers.DeleteExercise)
}