package main

import (
	_ "admin/docs"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
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
	app := fiber.New()

	app.Get("/swagger/*", swagger.WrapHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	CombineRoutes(app)

	app.Listen(":3000")	
}
