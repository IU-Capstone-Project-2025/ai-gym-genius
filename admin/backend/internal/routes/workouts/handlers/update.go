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

	if workoutUpdate.DurationNS != nil {
		workout.Duration = time.Duration(*workoutUpdate.DurationNS)
	}

	if workoutUpdate.StartTime != nil {
		workout.StartTime = *workoutUpdate.StartTime
	}

	// replace exercise sets
	if workoutUpdate.ExerciseSets != nil {
		exerciseSets := []schemas.ExerciseSet{}
		for _, exerciseSet := range *workoutUpdate.ExerciseSets {
			exerciseSet := &schemas.ExerciseSet{
				ExerciseID: exerciseSet.ExerciseID,
				Reps:       exerciseSet.Reps,
				Weight:     exerciseSet.Weight,
			}
			exerciseSets = append(exerciseSets, *exerciseSet)
		}

		err =
			database.DB.
				Unscoped().
				Model(workout).
				Association("ExerciseSets").
				Unscoped().
				Replace(exerciseSets)

		if err != nil {
			if errors.Is(err, gorm.ErrForeignKeyViolated) {
				return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
					Error: "invalid exercise_id provided",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
				Error: "failed to update exercise sets",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "Workout updated successfully",
	})
}

// AppendExerciseSet
// @Summary Append an exercise set to a workout
// @Tags workouts
// @Accept json
// @Produce json
// @Param ExerciseSet body models.ExerciseSetCreate true "ExerciseSet create payload"
// @Param id path int true "Workout ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Workout Not Found"
// @Failure 404 {object} models.ErrorResponse "Exercise Set Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/{id}/exercise_set [post]
func AppendExerciseSet(c *fiber.Ctx) error {
	workoutID, err := c.ParamsInt("id")

	if err != nil || workoutID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid workout id",
		})
	}

	exerciseSetCreate := &models.ExerciseSetCreate{}
	if err := c.BodyParser(exerciseSetCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	exerciseSet := &schemas.ExerciseSet{
		ExerciseID: exerciseSetCreate.ExerciseID,
		Reps:       exerciseSetCreate.Reps,
		Weight:     exerciseSetCreate.Weight,
	}

	err =
		database.DB.
			Model(&schemas.Workout{ID: uint(workoutID)}).
			Association("ExerciseSets").
			Append(exerciseSet)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) ||
			errors.Is(err, gorm.ErrForeignKeyViolated) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "no workout or exercise found",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Failed to append an exercise set to the workout",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "Exercise set appended successfully",
	})
}

// DeleteExerciseSet
// @Summary Delete an exercise set from a workout
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Param exercise_id query int true "Exercise ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Workout Not Found"
// @Failure 404 {object} models.ErrorResponse "Exercise Set Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/{id}/exercise_set [delete]
func DeleteExerciseSet(c *fiber.Ctx) error {
	workoutID, err := c.ParamsInt("id")
	if err != nil || workoutID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid workout id",
		})
	}

	exerciseID := c.QueryInt("exercise_id")
	if exerciseID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid or missing exercise_id query param",
		})
	}

	err =
		database.DB.
			Unscoped(). // perform hard-delete instead of soft-delete. wtf gorm??
			Model(&schemas.Workout{ID: uint(workoutID)}).
			Association("ExerciseSets").
			Unscoped().
			Delete(&schemas.ExerciseSet{
				WorkoutID: uint(workoutID), ExerciseID: uint(exerciseID),
			})

	// TODO add a hook (?? wtf gorm again) to check RowsAffected (below check doesnt work)
	// to see if workouts or exercise were not found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "no workout or exercise found",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "failed to delete an exercise set from the workout",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "exercise set deleted successfully",
	})
}
