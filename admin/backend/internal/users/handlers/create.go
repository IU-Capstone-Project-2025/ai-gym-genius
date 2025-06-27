package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"

	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	userCreate := &models.UserCreate{}
	
	if err := c.BodyParser(userCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	user := &schemas.User{
		Login: userCreate.Login,
		Hash:  database.Hash(userCreate.Login, userCreate.Password),
	}
	
	if err := database.DB.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"id":      user.ID,
	})
}
