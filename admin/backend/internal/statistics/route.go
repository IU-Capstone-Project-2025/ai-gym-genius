package statistics

import (
	"admin/internal/statistics/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatisticsRoutes(app *fiber.App) {
	app.Get("/statistics/active-users", handlers.GetNumberOfActiveUsers)
	app.Post("/statistics/add-activity", handlers.PostUserActivity)
}
