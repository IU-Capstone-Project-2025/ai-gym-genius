package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"github.com/gofiber/fiber/v2"
)

func DeleteWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "required 'id' parameter is malformed; should be a uint",
		})
	}
	
	result := database.DB.Delete(&schemas.Workout{}, id)
	
	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete workout",
		})
	}
	
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "workout not found",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "workout deleted successfully",
	})
}