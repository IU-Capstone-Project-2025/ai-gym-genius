package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"github.com/gofiber/fiber/v2"
	"time"
)

// CreateWorkout
// @Summary Create a new workout
// @Description Create a new workout with duration, start time, description, and weight
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body models.WorkoutCreate true "Workout create payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /workouts [post]
func CreateWorkout(c *fiber.Ctx) error {
	workoutCreate := &models.WorkoutCreate{}

	if err := c.BodyParser(workoutCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if workoutCreate.UserID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID must be greater than 0",
		})
	}

	if workoutCreate.Duration < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Duration must be greater than 0",
		})
	}

	if workoutCreate.Timestamp.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Timestamp must be provided",
		})
	}

		workout := &schemas.Workout{
		UserID:      workoutCreate.UserID,
		Duration:    workoutCreate.Duration,
		Timestamp:   workoutCreate.Timestamp,
	}

	user := &schemas.User{}
	if err := database.DB.First(user, workout.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Check if the user is active
	if user.Status != "active" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "User is not active",
		})
	}

	user.NumberOfWorkouts++
	user.TotalTimeSpent += workout.Duration
	user.AverageWorkoutDuration = user.TotalTimeSpent / time.Duration(user.NumberOfWorkouts)

	if err := database.DB.Save(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user workout count",
		})
	}

	if err := database.DB.Create(workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create workout",
		})
	}

	

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Workout created successfully",
		"id":      workout.ID,
	})
}
