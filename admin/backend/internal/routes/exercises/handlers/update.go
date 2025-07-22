package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateExercise
// @Security BearerAuth
// @Summary Update an existing exercise by ID
// @Tags exercises
// @Accept json
// @Produce json
// @Param exercise body models.ExerciseUpdate true "Exercise update payload"
// @Param id path int true "Exercise ID"
// @Success 200 {object} models.MessageResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Exercise Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises/{id} [patch]
func UpdateExercise(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "malformed id parameter; should be > 0",
		})
	}

	exerciseUpdate := &models.ExerciseUpdate{}

	if err := c.BodyParser(exerciseUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	exercise := &schemas.Exercise{}

	if err := database.DB.First(exercise, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "exercise not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to query exercise",
		})
	}

	// Update exercise fields only if they are provided in the input
	if exerciseUpdate.Name != nil {
		exercise.Name = *exerciseUpdate.Name
	}

	if exerciseUpdate.ImagePath != nil {
		exercise.ImagePath = *exerciseUpdate.ImagePath
	}

	if exerciseUpdate.Description != nil {
		exercise.Description = *exerciseUpdate.Description
	}

	if exerciseUpdate.MuscleGroups != nil {
		exercise.MuscleGroups = *exerciseUpdate.MuscleGroups
	}

	if err := database.DB.Save(&exercise).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to update exercise",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "exercise updated successfully",
	})
}
