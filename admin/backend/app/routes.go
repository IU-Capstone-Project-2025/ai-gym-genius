package main

import (
	"admin/internal/models"
	"admin/internal/routes/authorization"
	"admin/internal/routes/exercises"
	"admin/internal/routes/statistics"
	"admin/internal/routes/users"
	"admin/internal/routes/workouts"
	"embed"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

//go:embed welcome.html
var welcomePage embed.FS

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	statistics.SetupStatisticsRoutes(app)
	users.SetUpUserRoutes(app)
	workouts.SetUpWorkoutRoutes(app)
	exercises.SetupExerciseRoutes(app)

	app.Get("/swagger/*", swagger.WrapHandler)

	// healthcheck
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// welcome page
	app.Get("/", func(c *fiber.Ctx) error {
		html, err := welcomePage.ReadFile("welcome.html")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
				Error: "failed loading welcome page",
			})
		}
		c.Type("html")
		return c.Send(html)
	})
}
