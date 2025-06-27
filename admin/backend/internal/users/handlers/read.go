package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"

	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "'id' parameter is malformed; should be a uint",
		})
	}

	user := &schemas.User{}

	if err := database.DB.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve user",
		})
	}

	userRead := models.UserRead{
		ID:    user.ID,
		Login: user.Login,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user created successfully",
		"user": userRead,
	})
}
