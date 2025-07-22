package statistics

import (
	mw "admin/internal/middlewares"
	"admin/internal/routes/statistics/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatisticsRoutes(app *fiber.App) {
	app.Get("/statistics/active-users", mw.JWT, mw.AdminOnly, handlers.GetNumberOfActiveUsers)
	app.Get("/statistics/user/:id/", mw.JWT, mw.AdminOnly, handlers.GetUserActivityStats)
	app.Post("/statistics/add-activity", mw.JWT, handlers.PostUserActivity)
}
