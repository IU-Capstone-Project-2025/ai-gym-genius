package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"github.com/gofiber/fiber/v2"
)

// DeleteWorkout
// @Summary Delete a workout by ID
// @Description Delete a workout by their unique ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} map[string]string "Workout deleted successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Workout Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /workouts/{id} [delete]
func DeleteWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "required 'id' parameter is malformed; should be > 0",
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
