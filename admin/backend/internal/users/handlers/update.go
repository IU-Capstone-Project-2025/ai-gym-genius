package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateUser(c *fiber.Ctx) error {
	userUpdate := &models.UserUpdate{}
	
	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	
	user := &schemas.User{}
	
	if err := database.DB.First(user, userUpdate.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to query user",
		})
	}
	
	// TODO finish ths method
	// decide how to handle password updates
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "bleh",
	})
}
