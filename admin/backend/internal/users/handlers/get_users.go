package handlers

import (
	"admin/internal/database"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieves a list of all users from the database.
// @Tags         users
// @Produce      json
// @Success      200  {array}   schemas.User
// @Failure      500  {object}  string
// @Router       /users/all [get]
func GetAllUsers(c *fiber.Ctx) error {
	users, err := database.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.JSON(users)
}