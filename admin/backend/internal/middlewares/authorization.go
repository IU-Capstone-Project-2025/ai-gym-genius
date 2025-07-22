package middleware

import (
	"admin/config"
	"admin/internal/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type userIDKeyT struct{}

// key, under which current user's id is stored in c.Locals
var IDKey userIDKeyT

type userLoginKeyT struct{}

// key, under which current user's login is stored in c.Locals
var LoginKey userLoginKeyT

type userRoleKeyT struct{}

// key, under which current user's roles is stored in c.Locals
//
// admin or user
var RoleKey userRoleKeyT

func JWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Missing authorization header",
		})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Token not provided",
		})
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.C.JwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Invalid token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
				Error: "Token expired",
			})
		}
		userID, _ := claims["id"].(float64)
		login, _ := claims["login"].(string)
		role, _ := claims["role"].(string)
		c.Locals(IDKey, uint(userID))
		c.Locals(LoginKey, login)
		c.Locals(RoleKey, role)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Invalid token claims",
		})
	}

	return c.Next()
}
