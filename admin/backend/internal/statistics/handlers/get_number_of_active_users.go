package handlers

import (
	"github.com/gofiber/fiber/v2"
	"admin/internal/database/models"
)

func GetNumberOfActiveUsers(c *fiber.Ctx) error {
	data := new(models.GetNumberOfActiveUsersInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.StartDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'StartDate' cannot be empty",
		})
	}

	if data.EndDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'EndDate' cannot be empty",
		})
	}

	if data.Step == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Step' cannot be empty",
		})
	}

	//TODO add more validation for the date format and step value

	

	// This is a placeholder for the actual logic to get the number of active users.
	// In a real application, you would typically query a database or an external service.
	activeUsersCount := 42 // Example static value

	// Respond with the number of active users
	return c.JSON(fiber.Map{
		"active_users": activeUsersCount,
	})
}