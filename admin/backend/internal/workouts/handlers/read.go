package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetWorkout
// @Summary Get a workout by ID
// @Description Retrieve a workout by their unique ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} models.WorkoutRead
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Workout Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /workouts/{id} [get]
func GetWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "'id' parameter is malformed; should be > 0",
		})
	}

	workout := &schemas.Workout{}

	if err := database.DB.First(workout, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "workout not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve workout",
		})
	}

	workoutRead := models.WorkoutRead{
		Duration:    workout.Duration,
		StartTime:   workout.StartTime,
		Description: workout.Description,
		Weight:      workout.Weight,
	}

	return c.Status(fiber.StatusOK).JSON(workoutRead)
}
