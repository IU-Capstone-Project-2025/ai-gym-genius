package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
	
)

// AdminLoginHandler godoc
// @Summary Admin Login
// @Description Authenticate an admin with login and password, returning a token
// @Tags auth
// @Accept json
// @Produce json
// @Param authInput body models.AuthInput true "Login and Password"
// @Success 200 {object} models.TokenResponse "Admin logged in successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request body or missing fields"
// @Failure 404 {object} models.ErrorResponse "Admin not found"
// @Failure 401 {object} models.ErrorResponse "Incorrect password"
// @Failure 500 {object} models.ErrorResponse "Failed to query database or create token"
// @Router /auth_admin [post]
func AdminLoginHandler(c *fiber.Ctx) error {
	data := &models.AuthInput{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	if data.Login == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "field login cannot be empty",
		})
	}

	if data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "field password cannot be empty",
		})
	}

	var admin schemas.Admin
	if err := database.DB.Where("login = ?", data.Login).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to query database",
		})
	}

	hash := schemas.Hash(data.Login, data.Password)

	if admin.Hash != hash {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "incorrect password",
		})
	}

	token, err := admin.CreateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to create token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "Bearer " + token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,         
		Secure:   true,         
		SameSite: "Lax",        
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "logged in successfully",
	})
}
