package statistics

import (
	"admin/internal/routes/statistics/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatisticsRoutes(app *fiber.App) {
	app.Get("/statistics/active-users", handlers.GetNumberOfActiveUsers)
	app.Get("/statistics/user/:id/", handlers.GetUserActivityStats)
	app.Post("/statistics/add-activity", handlers.PostUserActivity)
}
