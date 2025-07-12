package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// LoginHandler godoc
// @Summary User Login
// @Description Authenticate a user with login and password, returning a token
// @Tags auth
// @Accept json
// @Produce json
// @Param authInput body models.AuthInput true "Login and Password"
// @Success 200 {object} map[string]string "User logged in successfully"
// @Failure 400 {object} map[string]string "Invalid request body or missing fields"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 401 {object} map[string]string "Incorrect password"
// @Failure 500 {object} map[string]string "Failed to query database or create token"
// @Router /auth [post]
func LoginHandler(c *fiber.Ctx) error {
	data := &models.AuthInput{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Login == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'login' cannot be empty",
		})
	}

	if data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'password' cannot be empty",
		})
	}

	var user schemas.Admin
	if err := database.DB.Where("Login = ?", data.Login).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to query database",
			})
		}
	}

	hash := database.Hash(data.Login, data.Password)

	if user.Hash != hash {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	token, err := database.CreateTokenForUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	c.Set("Authorization", "Bearer " + token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
	})
}
