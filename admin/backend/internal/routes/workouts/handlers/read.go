package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/middlewares"
	"admin/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetWorkout
// @Summary Get a workout by ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} models.WorkoutRead
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Workout Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/{id} [get]
func GetWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "id parameter is malformed; should be > 0",
		})
	}

	workout := &schemas.Workout{}

	if err := database.DB.Preload("ExerciseSets").First(workout, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "workout not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to retrieve workout",
		})
	}

	userIDRaw := c.Locals(middleware.IDKey)
	roleRaw := c.Locals(middleware.RoleKey)

	userID, ok := userIDRaw.(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (user ID)",
		})
	}

	role, ok := roleRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (role)",
		})
	}

	if int(userID) != int(workout.UserID) && role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "You can only get your own workouts",
		})
	}

	var exerciseSets []models.ExerciseSetRead
	for _, exerciseSet := range workout.ExerciseSets {
		exerciseSets = append(exerciseSets, models.ExerciseSetRead{
			WorkoutID:  exerciseSet.WorkoutID,
			ExerciseID: exerciseSet.ExerciseID,
			Weight:     exerciseSet.Weight,
			Reps:       exerciseSet.Reps,
		})
	}

	workoutRead := models.WorkoutRead{
		ID:           workout.ID,
		DurationNS:   workout.Duration.Nanoseconds(),
		Timestamp:    workout.StartTime,
		UserID:       workout.UserID,
		ExerciseSets: exerciseSets,
	}

	return c.Status(fiber.StatusOK).JSON(workoutRead)
}

// GetMyWorkouts
// @Summary Get user's workouts
// @Tags workouts
// @Accept json
// @Produce json
// @Success 200 {object} []models.WorkoutRead
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/my [get]
func GetMyWorkouts(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.IDKey).(float64)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "malformed claims or unauthorized",
		})
	}

	var workouts []schemas.Workout

	if err := database.DB.Preload("ExerciseSets").Find(&workouts, "user_id = ?", uint(userID)).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to fetch user's workouts",
		})
	}

	var workoutReads []models.WorkoutRead
	for _, workout := range workouts {
		var exerciseSetReads []models.ExerciseSetRead
		for _, exerciseSet := range workout.ExerciseSets {
			exerciseSetReads = append(exerciseSetReads, models.ExerciseSetRead{
				Weight:     exerciseSet.Weight,
				Reps:       exerciseSet.Reps,
				ExerciseID: exerciseSet.ExerciseID,
				WorkoutID:  exerciseSet.WorkoutID,
			})
		}
		workoutReads = append(workoutReads, models.WorkoutRead{
			ID:           workout.ID,
			DurationNS:   workout.Duration.Nanoseconds(),
			Timestamp:    workout.StartTime,
			ExerciseSets: exerciseSetReads,
		})
	}

	return c.Status(fiber.StatusOK).JSON(workoutReads)
}
