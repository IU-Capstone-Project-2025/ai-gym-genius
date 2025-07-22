package users

import (
	middleware "admin/internal/middlewares"
	"admin/internal/routes/users/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Post("/users/", handlers.CreateUser)
	app.Get("/users/", middleware.JWT, middleware.AdminOnly, handlers.GetUsersPaginate)
	app.Get("/users/count", middleware.JWT, middleware.AdminOnly, handlers.GetUserCount)
	app.Get("/users/me", middleware.JWT, handlers.GetCurrentUser)
	app.Get("/users/:id", middleware.JWT, middleware.AdminOnly, handlers.GetUserByID)
	app.Patch("/users/me", middleware.JWT, handlers.UpdateCurrentUser)
	app.Patch("/users/:id", middleware.JWT, middleware.AdminOnly, handlers.UpdateUserByID)
	app.Delete("/users/me", middleware.JWT, handlers.DeleteCurrentUser)
	app.Delete("/users/:id", middleware.JWT, middleware.AdminOnly, handlers.DeleteUserByID)
}
