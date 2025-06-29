package handlers

import(
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"github.com/gofiber/fiber/v2"
)

func CreateWorkout(c *fiber.Ctx) error {
	workoutCreate := &models.WorkoutCreate{}

	if err := c.BodyParser(workoutCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	workout := &schemas.Workout{
		Duration:    workoutCreate.Duration,
		StartTime:   workoutCreate.StartTime,
		Description: workoutCreate.Description,
		Weight:      workoutCreate.Weight,
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