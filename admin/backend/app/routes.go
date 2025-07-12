package main

import (
	"admin/internal/routes/authorization"
	"admin/internal/routes/exercises"
	"admin/internal/routes/statistics"
	"admin/internal/routes/users"
	"admin/internal/routes/workouts"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	statistics.SetupStatisticsRoutes(app)
	users.SetUpUserRoutes(app)
	workouts.SetUpWorkoutRoutes(app)
	exercises.SetupExerciseRoutes(app)
	
	app.Get("/swagger/*", swagger.WrapHandler)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
