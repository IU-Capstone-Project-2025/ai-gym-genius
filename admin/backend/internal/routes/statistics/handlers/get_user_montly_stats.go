package handlers

import (
	"admin/internal/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetUserActivityStats
// @Summary Get user activity statistics
// @Description Retrieve monthly activity statistics for a user
// @Tags statistics
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} map[string]interface{} "Monthly activity statistics"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id}/activity [get]
func GetUserActivityStats(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil || userID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	stats, err := GetUserMonthlyStats(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not fetch stats",
		})
	}

	result := []fiber.Map{}
	for _, stat := range stats {
		result = append(result, fiber.Map{
			"month": stat.Month.Format("2006-01"),
			"count": stat.Count,
		})
	}

	return c.JSON(result)
}

type MonthlyStat struct {
	Month time.Time
	Count int
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