package users

import (
	"admin/internal/routes/users/handlers"
	"admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Post("/users/", handlers.CreateUser)
	app.Get("/users/", middleware.JWTMiddleware, handlers.GetUsersPaginate)
	app.Get("/users/count", middleware.JWTMiddleware, handlers.GetUserCount)
	app.Get("/users/:id", middleware.JWTMiddleware, handlers.GetUser)
	app.Patch("/users/:id", middleware.JWTMiddleware, handlers.UpdateUser)
	app.Delete("/users/:id", middleware.JWTMiddleware, handlers.DeleteUser)
}
