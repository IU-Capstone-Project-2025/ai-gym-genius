package main

import (
	_ "admin/docs"
	"admin/internal/database"
	middleware "admin/internal/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Gym Genius API
// @version 1.0
// @description API for Gym Genius application
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	if err := database.InitDatabase(); err != nil {
		panic(err) // failed to connect or migrate
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))

	app.Use(middleware.LoggingMiddleware())

	CombineRoutes(app)
	
	app.Use(func(c *fiber.Ctx) error { // 404 handler
		return c.SendStatus(fiber.StatusNotFound)
	})
	
	log.Fatal(app.Listen(":3000"))
}
