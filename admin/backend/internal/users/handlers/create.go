package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateUser
// @Summary Create a new user
// @Description Create a new user with login and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "User create payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "user with this login already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"id":      user.ID,
	})
}
