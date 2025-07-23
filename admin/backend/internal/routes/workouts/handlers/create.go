package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"time"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	middleware "admin/internal/middlewares"
)

// CreateWorkout
// @Summary Create a new workout
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body models.WorkoutCreate true "Workout create payload"
// @Success 200 {object} models.CreatedResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts [post]
func CreateWorkout(c *fiber.Ctx) error {
	workoutCreate := &models.WorkoutCreate{}

	if err := c.BodyParser(workoutCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	userIDRaw := c.Locals(middleware.IDKey)

	userID, ok := userIDRaw.(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (user ID)",
		})
	}

	if int(userID) != int(workoutCreate.UserID) {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "You can only create your own workouts",
		})
	}

	if workoutCreate.UserID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "user id must be greater than 0",
		})
	}

	workoutDuration := time.Duration(workoutCreate.DurationNS)

	if workoutDuration < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "duration must be greater than 0",
		})
	}

	if workoutCreate.StartTime.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "timestamp must be provided",
		})
	}

	var exerciseSets []schemas.ExerciseSet
	for _, exerciseSet := range workoutCreate.ExerciseSets {
		exerciseSets = append(exerciseSets, schemas.ExerciseSet{
			ExerciseID: exerciseSet.ExerciseID,
			Weight:     exerciseSet.Weight,
			Reps:       exerciseSet.Reps,
		})
	}

	workout := &schemas.Workout{
		UserID:       workoutCreate.UserID,
		Duration:     workoutDuration,
		StartTime:    workoutCreate.StartTime,
		ExerciseSets: exerciseSets,
	}

	if err := database.DB.Create(workout).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "user for this workout not found",
			})
		} else if errors.Is(err, schemas.ErrUserNotActive) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Error: "user for this workout is not active",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
				Error: "failed to create workout",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.CreatedResponse{
		Message: "workout created successfully",
		ID:      workout.ID,
	})
}
