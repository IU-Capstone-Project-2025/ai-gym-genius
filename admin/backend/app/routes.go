package main

import (
	"admin/internal/authorization"
	"admin/internal/statistics"
	swagger "github.com/swaggo/fiber-swagger"
	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	statistics.SetupStatisticsRoutes(app)
	app.Get("/swagger/*", swagger.WrapHandler)
}
