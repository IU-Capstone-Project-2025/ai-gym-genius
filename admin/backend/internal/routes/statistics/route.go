package statistics

import (
	middleware "admin/internal/middlewares"
	"admin/internal/routes/statistics/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupStatisticsRoutes(app *fiber.App) {
	app.Get("/statistics/active-users", middleware.JWTMiddleware, handlers.GetNumberOfActiveUsers)
	app.Get("/statistics/user/:id/", middleware.JWTMiddleware, handlers.GetUserActivityStats)
	app.Post("/statistics/add-activity", middleware.JWTMiddleware, handlers.PostUserActivity)
}
