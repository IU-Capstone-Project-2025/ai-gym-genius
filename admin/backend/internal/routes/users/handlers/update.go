package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	"time"
	"admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateUserByID
// @Security BearerAuth
// @Summary Update any user by ID (Admin only)
// @Description Update any user by their ID (Admin privileges required)
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserUpdate true "User update payload"
// @Success 200 {object} models.MessageResponse "Updated successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Forbidden"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/{id} [patch]
func UpdateUserByID(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil || id < 1 {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "invalid user id",
        })
    }

    roleRaw := c.Locals(middleware.RoleKey)
    role, ok := roleRaw.(string)
    if !ok || role != "admin" {
        return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
            Error: "Admin privileges required",
        })
    }

    input := &models.UserUpdate{}
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "invalid request body",
        })
    }

    return updateUserByID(uint(id), input, c)
}

// UpdateCurrentUser
// @Security BearerAuth
// @Summary Update current user
// @Description Update the currently authenticated user's data
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.UserUpdate true "User update payload"
// @Success 200 {object} models.MessageResponse "Updated successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/me [patch]
func UpdateCurrentUser(c *fiber.Ctx) error {
    userIDRaw := c.Locals(middleware.IDKey)
    
    userID, ok := userIDRaw.(uint)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
            Error: "Unauthorized or invalid token (user ID)",
        })
    }

    input := &models.UserUpdate{}
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "invalid request body",
        })
    }

    return updateUserByID(userID, input, c)
}

func updateUserByID(id uint, input *models.UserUpdate, c *fiber.Ctx) error {
    user := &schemas.User{}
    if err := database.DB.First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
                Error: "user not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
            Error: "failed to query user",
        })
    }

    // Update user fields only if they are provided in the input
    if input.Login != nil {
        user.Login = *input.Login
    }

    if input.Name != nil {
        user.Name = *input.Name
    }

    if input.Surname != nil {
        user.Surname = *input.Surname
    }

    if input.Email != nil {
        user.Email = *input.Email
    }

    if input.SubscriptionPlan != nil {
        user.SubscriptionPlan = *input.SubscriptionPlan
    }

    if input.Status != nil {
        user.Status = *input.Status
    }

    if input.LastActivity != nil {
        user.LastActivity = *input.LastActivity
    }

    if input.NumberOfWorkouts != nil {
        user.NumberOfWorkouts = *input.NumberOfWorkouts
    }

    if input.TotalTimeSpentNS != nil {
        user.TotalTimeSpent = time.Duration(*input.TotalTimeSpentNS)
    }

    if input.StreakCount != nil {
        user.StreakCount = *input.StreakCount
    }

    if input.AverageWorkoutDurationNS != nil {
        user.AverageWorkoutDuration = time.Duration(*input.AverageWorkoutDurationNS)
    }

    if input.Password != nil {
        hashed := schemas.Hash(*input.Login, *input.Password)
        user.Hash = hashed
    }

    if err := database.DB.Save(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
            Error: "failed to update user",
        })
    }

    return c.Status(fiber.StatusOK).JSON(models.MessageResponse{
		Message: "user updated successfully",
	})
}