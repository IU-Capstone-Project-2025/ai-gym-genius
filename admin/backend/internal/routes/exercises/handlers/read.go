package handlers

import (
	"admin/internal/database"
	"admin/internal/models"
	"admin/internal/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetExerciseByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid exercise ID",
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
			"error": "failed to retrieve exercise",
		})
	}

	exerciseRead := models.ExerciseRead{
		ID:          exercise.ID,
		Name:        exercise.Name,
		Description: exercise.Description,
		MuscleGroup: exercise.MuscleGroup,
		URL: 	   exercise.URL,

	}

	return c.Status(fiber.StatusOK).JSON(exerciseRead)
}

func GetExercisesPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}
	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var exercises []schemas.Exercise
	if err := database.DB.Limit(limit).Offset(offset).Find(&exercises).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve users",
		})
	}

	exerciseReads := make([]models.ExerciseRead, len(exercises))
	for i, exercise := range exercises {
		exerciseReads[i] = models.ExerciseRead{
			ID:    exercise.ID,
			Name: exercise.Name,
			Description: exercise.Description,
			MuscleGroup: exercise.MuscleGroup,
			URL: exercise.URL,
		}
	}

	return c.Status(fiber.StatusOK).JSON(exerciseReads)
}