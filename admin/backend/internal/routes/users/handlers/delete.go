package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	middleware "admin/internal/middlewares"
	"admin/internal/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteUserByID
// @Security BearerAuth
// @Summary Delete any user by ID (Admin only)
// @Description Delete any user by their ID (Admin privileges required)
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
func DeleteUserByID(c *fiber.Ctx) error {
    targetID, err := c.ParamsInt("id")
    if err != nil || targetID <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "'id' parameter is malformed; should be > 0",
        })
    }
    return deleteUserByID(uint(targetID), c)
}

// DeleteCurrentUser
// @Security BearerAuth
// @Summary Delete current user
// @Description Delete the currently authenticated user's account
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.MessageResponse "User deleted successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/me [delete]
func DeleteCurrentUser(c *fiber.Ctx) error {
    userID := c.Locals(middleware.IDKey).(uint)
    return deleteUserByID(userID, c)
}

func deleteUserByID(id uint, c *fiber.Ctx) error {
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