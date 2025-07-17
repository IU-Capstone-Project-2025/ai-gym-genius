package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	middleware "admin/internal/middlewares"
	"admin/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser
// @Security BearerAuth
// @Summary Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.MessageResponse "User deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Forbidden"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	targetIDParam := c.Params("id")
	
	targetID, err := strconv.Atoi(targetIDParam)
	if err != nil || targetID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "Invalid user ID in path",
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

	if int(userID) != targetID && role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "You can only delete your own account",
		})
	}

	result := database.DB.Delete(&schemas.User{}, targetID)

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "Failed to delete user",
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "User deleted successfully",
	})
}
