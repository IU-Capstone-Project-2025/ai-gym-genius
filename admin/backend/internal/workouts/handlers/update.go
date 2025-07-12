package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateWorkout
// @Summary Update an existing workout
// @Description Update workout details by ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body models.WorkoutUpdate true "Workout update payload"
// @Param id path int true "Workout ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Workout Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /workouts/{id} [patch]
func UpdateWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed 'id' parameter; should be > 0",
		})
	}
	
	workoutUpdate := &models.WorkoutUpdate{}
	
	if err := c.BodyParser(workoutUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
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
			"error": "failed to query workout",
		})
	}

	// Update workout fields only if they are provided in the input
	if workoutUpdate.UserID != nil {
		workout.UserID = *workoutUpdate.UserID
	}

	if workoutUpdate.Duration != nil {
		workout.Duration = *workoutUpdate.Duration
	}

	if workoutUpdate.Timestamp != nil {
		workout.Timestamp = *workoutUpdate.Timestamp
	}

	if err := database.DB.Save(&workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update workout",
		})
	}

	return c.JSON(workout)
}
