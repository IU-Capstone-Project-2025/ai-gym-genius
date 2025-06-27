package main

import (
	"admin/internal/authorization"
	"admin/internal/statistics"
	"admin/internal/users"
	
	swagger "github.com/swaggo/fiber-swagger"
	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	statistics.SetupStatisticsRoutes(app)
	users.SetUpUserRoutes(app)
	
	app.Get("/swagger/*", swagger.WrapHandler)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
