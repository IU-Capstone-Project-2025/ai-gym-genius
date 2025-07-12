package exercises

import (
	"admin/internal/routes/exercises/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupExerciseRoutes(app *fiber.App) {
	app.Post("/exercises/", handlers.AddExercise)
	app.Get("/exercises/", handlers.GetExercisesPaginate)
	app.Get("/exercises/:id", handlers.GetExerciseByID)
	app.Patch("/exercises/:id", handlers.UpdateExercise)
	app.Delete("/exercises/:id", handlers.DeleteExercise)
}