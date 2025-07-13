package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateWorkout
// @Summary Update an existing workout by ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body models.WorkoutUpdate true "Workout update payload"
// @Param id path int true "Workout ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Workout Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/{id} [patch]
func UpdateWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "malformed 'id' parameter; should be > 0",
		})
	}

	workoutUpdate := &models.WorkoutUpdate{}

	if err := c.BodyParser(workoutUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	workout := &schemas.Workout{}

	if err := database.DB.First(workout, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "workout not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to query workout",
		})
	}

	// Update workout fields only if they are provided in the input
	if workoutUpdate.UserID != nil {
		workout.UserID = *workoutUpdate.UserID
	}

	workoutDuration := time.Duration(*workoutUpdate.DurationNS)

	if workoutUpdate.DurationNS != nil {
		workout.Duration = workoutDuration
	}

	if workoutUpdate.StartTime != nil {
		workout.StartTime = *workoutUpdate.StartTime
	}

	if workoutUpdate.ExerciseSets != nil {
		var exerciseSets []schemas.ExerciseSet
		for _, exerciseSet := range workout.ExerciseSets {
			exerciseSets = append(exerciseSets, schemas.ExerciseSet{
				WorkoutID:  exerciseSet.WorkoutID,
				ExerciseID: exerciseSet.ExerciseID,
				Weight:     exerciseSet.Weight,
				Reps:       exerciseSet.Reps,
			})
		}
		workout.ExerciseSets = exerciseSets
	}

	if err := database.DB.Save(workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to update workout",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "Workout updated successfully",
	})
}
