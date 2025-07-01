package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateUser
// @Summary Update an existing user
// @Description Update user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserUpdate true "User update payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "User Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed 'id' parameter; should be > 0",
		})
	}
	
	userUpdate := &models.UserUpdate{}

	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
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
			"error": "failed to query user",
		})
	}

	// TODO finish ths method
	// decide how to handle password updates
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "bleh",
	})
}
