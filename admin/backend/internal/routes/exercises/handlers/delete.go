package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

// DeleteExercise
// @Summary Delete an exercise by id
// @Tags exercises
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} models.MessageResponse "Exercise deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Exercise Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises/{id} [delete]
func DeleteExercise(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "required id parameter is malformed; should be > 0",
		})
	}

	roleRaw := c.Locals(middleware.RoleKey)

	role, ok := roleRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (role)",
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "This endpoint is restricted to admin users only",
		})
	}

	result := database.DB.Delete(&schemas.Exercise{}, id)

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to delete exercise",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: "exercise not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "exercise deleted successfully",
	})
}
