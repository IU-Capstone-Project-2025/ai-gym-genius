package statistics

import (
	"github.com/gofiber/fiber/v2"
	"admin/internal/statistics/handlers"
)

func SetupStatisticsRoutes(app *fiber.App) {
	app.Get("/statistics", handlers.GetNumberOfActiveUsers)
	app.Post("/statistics/add-activity", handlers.GetUserActivity)
}