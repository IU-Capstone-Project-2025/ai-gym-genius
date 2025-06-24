package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetNumberOfActiveUsers
// @Summary Get number of active users
// @Description Returns the count of active users within the specified date range and interval step (hour or day)
// @Tags Statistics
// @Accept json
// @Produce json
// @Param start_date query string true "Start date in YYYY-MM-DD format" example("2025-06-01")
// @Param end_date query string true "End date in YYYY-MM-DD format" example("2025-06-10")
// @Param step query string true "Interval step" Enums(hour, day) example("day")
// @Success 200 {object} map[string]int "Active users count"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Router /statistics [get]
func GetNumberOfActiveUsers(c *fiber.Ctx) error {
    startDate := c.Query("start_date")
    endDate := c.Query("end_date")
    step := c.Query("step")

    if startDate == "" || endDate == "" || step == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required query parameters"})
    }

	user_number := 42

    return c.JSON(fiber.Map{
        "active_users": user_number,
    })
}
