package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"errors"
	"admin/internal/middlewares"
)

// DeleteWorkout
// @Summary Delete a workout by ID
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} models.MessageResponse "Workout deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Workout Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /workouts/{id} [delete]
func DeleteWorkout(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "required 'id' parameter is malformed; should be > 0",
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
			Error: "You can only delete your own workouts",
		})
	}

	result := database.DB.Select("ExerciseSets").Delete(&schemas.Workout{}, id)

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to delete workout",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: "workout not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "workout deleted successfully",
	})
}
