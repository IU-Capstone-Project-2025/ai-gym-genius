package main

import (
	"admin/internal/statistics"

	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	statistics.SetupStatisticsRoutes(app)
}
