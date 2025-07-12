package handlers

import (
	"admin/internal/database"
	"admin/internal/database/models"
	"admin/internal/database/schemas"

	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUser
// @Summary Get a user by ID
// @Description Retrieve a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserRead
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "User Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "'id' parameter is malformed; should be > 0",
		})
	}

	user := &schemas.User{}

	if err := database.DB.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve user",
		})
	}

	userRead := models.UserRead{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
		Surname: user.Surname,
		Email: user.Email,
		SubscriptionType: user.SubscriptionType,
		Status: user.Status,
		LastActivity: user.LastActivity,
		NumberOfWorkouts: user.NumberOfWorkouts,
		TotalTimeSpent: user.TotalTimeSpent,
		StreakCount: user.StreakCount,
		AverageWorkoutDuration: user.AverageWorkoutDuration,
	}

	return c.Status(fiber.StatusOK).JSON(userRead)
}

// GetUserPaginate
// @Summary Get paginated list of users
// @Description Retrieve a paginated list of users with optional page and limit query parameters
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Success 200 {array} models.UserRead
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [get]
func GetUsersPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}
	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var users []schemas.User
	if err := database.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve users",
		})
	}

	userReads := make([]models.UserRead, len(users))
	for i, user := range users {
		userReads[i] = models.UserRead{
			ID:    user.ID,
			Login: user.Login,
			Name:  user.Name,
			Surname: user.Surname,
			Email: user.Email,
			SubscriptionType: user.SubscriptionType,
			Status: user.Status,
			LastActivity: user.LastActivity,
			NumberOfWorkouts: user.NumberOfWorkouts,
			TotalTimeSpent: user.TotalTimeSpent,
			StreakCount: user.StreakCount,
			AverageWorkoutDuration: user.AverageWorkoutDuration,
		}
	}

	return c.Status(fiber.StatusOK).JSON(userReads)
}
