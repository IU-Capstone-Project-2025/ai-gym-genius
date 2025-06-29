package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateWorkout(c *fiber.Ctx) error {
	workoutUpdate := &models.WorkoutUpdate{}
	
	if err := c.BodyParser(workoutUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	
	workout := &schemas.Workout{}
	
	if err := database.DB.First(workout, workoutUpdate.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "workout not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to query workout",
		})
	}

	if err := database.DB.Delete(workout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update old workout",
		})
	}

	updatedWorkout := &schemas.Workout{
		Duration:    workoutUpdate.Duration,
		StartTime:   workoutUpdate.StartTime,
		Description: workoutUpdate.Description,
		Weight:      workoutUpdate.Weight,
	}

	if err := database.DB.Create(updatedWorkout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update workout",
		})
	}

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "workout updated successfully",
	})
}
