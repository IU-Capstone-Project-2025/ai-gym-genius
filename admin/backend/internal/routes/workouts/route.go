package workouts

import (
	middleware "admin/internal/middlewares"
	"admin/internal/routes/workouts/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpWorkoutRoutes(app *fiber.App) {
	app.Post("/workouts/", middleware.JWT, handlers.CreateWorkout)
	app.Get("/workouts/my", middleware.JWT, handlers.GetMyWorkouts)
	app.Get("/workouts/:id", middleware.JWT, handlers.GetWorkout)
	app.Patch("/workouts/:id", middleware.JWT, handlers.UpdateWorkout)
	app.Delete("/workouts/:id", middleware.JWT, handlers.DeleteWorkout)
	app.Post("/workouts/:id/exercise_set", middleware.JWT, handlers.AppendExerciseSet)
	app.Delete("/workouts/:id/exercise_set", middleware.JWT, handlers.DeleteExerciseSet)
}
