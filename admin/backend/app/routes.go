package main

import (
	"admin/internal/authorization"
	"admin/internal/statistics"

	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	statistics.SetupStatisticsRoutes(app)
}
