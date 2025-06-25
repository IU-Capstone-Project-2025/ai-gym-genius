package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"

	"github.com/gofiber/fiber/v2"
)

// PostUserActivity
// @Summary Record user activity
// @Description Records activity for a user on a specific date
// @Tags Statistics
// @Accept json
// @Produce json
// @Param input body models.GetUserActivityInput true "Request parameters"
// @Success 200 {object} string "Success message"
// @Failure 400 {object} string "Validation error"
// @Failure 500 {object} string "Server error"
// @Router /statistics/add-activity [post]
func PostUserActivity(c *fiber.Ctx) error {
	data := new(models.GetUserActivityInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'user_id' cannot be empty",
		})
	}

	if data.Date.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'date' cannot be empty",
		})
	}

	record := schemas.UserActivity{
		UserID: data.UserID,
		Date:   data.Date,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to record user activity",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User activity recorded successfully",
		"id": record.ID,
	})

}
