package handlers

import (
	"github.com/gofiber/fiber/v2"
	"admin/internal/database/models"
	"admin/internal/database"
)

func GetUserActivity(c *fiber.Ctx) error {
	data := new(models.GetUserActivityInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'UserID' cannot be empty",
		})
	}

	if data.Date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Date' cannot be empty",
		})
	}

	//TODO add more validation for the date format

	record := models.UserActivity{
		UserID:  data.UserID,
		Date: data.Date,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to record user activity",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User activity retrieved successfully",
	})

}