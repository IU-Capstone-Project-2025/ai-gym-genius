package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateExercise
// @Summary Update an existing exercise
// @Description Update exercise details by ID
// @Tags exercises
// @Accept json
// @Produce json
// @Param exercise body models.ExerciseUpdate true "Exercise update payload"
// @Param id path int true "Exercise ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Exercise Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /exercises/{id} [patch]
func UpdateExercise(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "malformed 'id' parameter; should be > 0",
		})
	}
	
	exerciseUpdate := &models.ExerciseUpdate{}
	
	if err := c.BodyParser(exerciseUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	
	exercise := &schemas.Exercise{}
	
	if err := database.DB.First(exercise, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "exercise not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to query exercise",
		})
	}

	updatedExercise := &schemas.Exercise{
		ID: uint(id),
		Name: exercise.Name,
		URL: exercise.URL,
	}

	if err := database.DB.Save(updatedExercise).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update exercise",
		})
	}

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "exercise updated successfully",
	})
}
