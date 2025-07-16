package workouts

import (
	middleware "admin/internal/middlewares"
	"admin/internal/routes/workouts/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpWorkoutRoutes(app *fiber.App) {
	app.Post("/workouts/", middleware.JWTMiddleware, handlers.CreateWorkout)
	app.Get("/workouts/:id", middleware.JWTMiddleware, handlers.GetWorkout)
	app.Patch("/workouts/:id", middleware.JWTMiddleware, handlers.UpdateWorkout)
	app.Delete("/workouts/:id", middleware.JWTMiddleware, handlers.DeleteWorkout)
	app.Post("/workouts/:id/exercise_set", middleware.JWTMiddleware, handlers.AppendExerciseSet)
	app.Delete("/workouts/:id/exercise_set", middleware.JWTMiddleware, handlers.DeleteExerciseSet)
}
