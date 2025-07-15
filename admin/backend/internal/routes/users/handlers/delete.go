package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	middleware "admin/internal/middlewares"
	"admin/internal/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser
// @Summary Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.MessageResponse "User deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	middleware.JWTMiddleware(c)
	id := c.Locals(middleware.IDKey).(int64)

	if id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "required id parameter is malformed; should be > 0",
		})
	}

	result := database.DB.Delete(&schemas.User{}, id)

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to delete user",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "user deleted successfully",
	})
}
