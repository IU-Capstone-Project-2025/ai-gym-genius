package handlers

import (
	"admin/internal/database"
	"admin/internal/models"
	"time"
	"admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

// GetUserActivityStats
// @Security BearerAuth
// @Summary Get user activity statistics
// @Description Retrieve monthly activity statistics for a user
// @Tags statistics
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []MonthlyStat "Monthly activity statistics"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/{id}/activity [get]
func GetUserActivityStats(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")

	if err != nil || userID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid user id",
		})
	}

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

	montlyStats, err := GetUserMonthlyStats(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "could not fetch stats",
		})
	}

	return c.Status(fiber.StatusOK).JSON(montlyStats)
}

type MonthlyStat struct {
	Month time.Time `json:"month"`
	Count int       `json:"count"`
}

func GetUserMonthlyStats(userID uint) ([]MonthlyStat, error) {
	var stats []MonthlyStat

	err := database.DB.
		Table("workouts").
		Select("DATE_TRUNC('month', timestamp) as month, COUNT(*) as count").
		Where("user_id = ?", userID).
		Group("month").
		Order("month").
		Scan(&stats).Error

	return stats, err
}
