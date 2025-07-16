package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	middleware "admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

// PostUserActivity
// @Summary Record user activity
// @Description Records activity for a user on a specific date
// @Tags statistics
// @Accept json
// @Produce json
// @Param input body models.UserActivityCreate true "Request parameters"
// @Success 200 {object} models.CreatedResponse "Success message"
// @Failure 400 {object} models.ErrorResponse "Validation error"
// @Failure 500 {object} models.ErrorResponse "Server error"
// @Router /statistics/add-activity [post]
func PostUserActivity(c *fiber.Ctx) error {
	roleRaw := c.Locals(middleware.RoleKey)

	role, ok := roleRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (role)",
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "This endpoint is restricted to admin users only",
		})
	}

	data := &models.UserActivityCreate{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if data.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Field 'user_id' cannot be empty",
		})
	}

	if data.Date.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Field 'date' cannot be empty",
		})
	}

	record := schemas.UserActivity{
		UserID: data.UserID,
		Date:   data.Date,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "Failed to record user activity",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.CreatedResponse{
		Message: "User activity recorded successfully",
		ID:      record.ID,
	})
}
