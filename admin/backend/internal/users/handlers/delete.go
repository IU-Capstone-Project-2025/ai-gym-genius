package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
		
	"github.com/gofiber/fiber/v2"
)

// DeleteUser
// @Summary Delete a user by ID
// @Description Delete a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string "User deleted successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "User Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "required 'id' parameter is malformed; should be > 0",
		})
	}
	
	result := database.DB.Delete(&schemas.User{}, id)
	
	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}
	
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
