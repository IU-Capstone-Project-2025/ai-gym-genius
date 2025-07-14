package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"

	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
    ErrInvalidStepParam = errors.New("malformed step parameter. should be of format '<uint><h | d>'");
)

type params struct {
    StartDate time.Time `query:"start_date"`
    EndDate   time.Time `query:"end_date"`
    Step      string    `query:"step"`
}

type step struct {
    unit  byte
    value uint64
}

type intervalResult struct {
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Count     int64     `json:"count"`
}

func parseStep(stepStr string) (step, error) {
    if len(stepStr) < 2 {
        return step{}, ErrInvalidStepParam
    }
    
    unit := stepStr[len(stepStr)-1]
    value, err := strconv.ParseUint(stepStr[:len(stepStr)-1], 10, 64)
    if (unit != 'h' && unit != 'd') || err != nil {
        return step{}, ErrInvalidStepParam
    }
    return step{unit: unit, value: value}, nil
}

// GetNumberOfActiveUsers
// @Summary Get number of active users in time intervals
// @Description Returns the count of distinct active users grouped by specified time intervals between start and end dates.
// @Tags statistics
// @Accept json
// @Produce json
// @Param start_date query string true "Start date/time in RFC3339 format (e.g., 2025-06-01T00:00:00Z)"
// @Param end_date query string true "End date/time in RFC3339 format (e.g., 2025-06-10T00:00:00Z)"
// @Param step query string true "Time interval step duration, format: <number><unit> where unit is 'h' for hours or 'd' for days (e.g., '24h', '7d')" example("24h")
// @Success 200 {object} []intervalResult "Array of interval results containing start time, end time, and user count"
// @Failure 400 {object} models.ErrorResponse "Invalid input parameters (missing, malformed, or invalid date range)"
// @Failure 500 {object} models.ErrorResponse "Internal server error while processing the request"
// @Router /statistics/active-users [get]
func GetNumberOfActiveUsers(c *fiber.Ctx) error {
    params := params{}
    if err := c.QueryParser(&params); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "missing or malformed query parameters",
        })
    }

    if params.StartDate.IsZero() || params.EndDate.IsZero() {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "both start_date and end_date are required",
        })
    }

    if params.EndDate.Before(params.StartDate) {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "end_date must be after start_date",
        })
    }

    step, err := parseStep(params.Step)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: err.Error(),
        })
    }

    var intervalDuration time.Duration
    switch step.unit {
    case 'h':
        intervalDuration = time.Duration(step.value) * time.Hour
    case 'd':
        intervalDuration = time.Duration(step.value) * 24 * time.Hour
    default:
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "invalid step unit, must be 'h' or 'd'",
        })
    }

    var intervals = []intervalResult{}
    currentStart := params.StartDate
    for currentStart.Before(params.EndDate) {
        currentEnd := currentStart.Add(intervalDuration)
        if currentEnd.After(params.EndDate) {
            currentEnd = params.EndDate
        }

        intervals = append(intervals, intervalResult{
            StartTime: currentStart,
            EndTime:   currentEnd,
        })

        currentStart = currentEnd
    }

    for i := range intervals {
        var count int64
        err :=
            database.DB.Model(&schemas.UserActivity{}).
            Where("date BETWEEN ? AND ?", intervals[i].StartTime, intervals[i].EndTime).
            Count(&count).Error

        if err != nil && err != gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
                Error: "failed to query database",
            })
        }

        intervals[i].Count = count
    }
    
    return c.Status(fiber.StatusOK).JSON(intervals)
}