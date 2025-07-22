package exercises

import (
	mw "admin/internal/middlewares"
	"admin/internal/routes/exercises/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupExerciseRoutes(app *fiber.App) {
	app.Post("/exercises/", mw.JWT, mw.AdminOnly, handlers.AddExercise)
	app.Get("/exercises/", mw.JWT, handlers.GetExercisesPaginate)
	app.Get("/exercises/photo/:id", mw.JWT, mw.AdminOnly, handlers.GetExercisePhoto)
	app.Get("/exercises/photo/", mw.JWT, handlers.GetAllExercisePhotos)
	app.Get("/exercises/:id", mw.JWT, handlers.GetExerciseByID)
	app.Patch("/exercises/:id", mw.JWT, mw.AdminOnly, handlers.UpdateExercise)
	app.Delete("/exercises/:id", mw.JWT, mw.AdminOnly, handlers.DeleteExercise)
}
