package middleware

import (
	"admin/config"
	"admin/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = config.C.JwtSecret

type userIDKeyT struct{}

var IDKey userIDKeyT

type userLoginKeyT struct{}

var LoginKey userLoginKeyT

type userRoleKeyT struct{}

var RoleKey userRoleKeyT

func JWT(c *fiber.Ctx) error {
	authHeader := c.Cookies("jwt")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Missing authorization header",
		})
	}

	// tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// if tokenString == "" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
	// 		Error: "Token not provided",
	// 	})
	// }

	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(jwtSecret), nil
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

		c.Locals(IDKey, claims["id"])
		c.Locals(LoginKey, claims["login"])
		c.Locals(RoleKey, claims["role"])
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Invalid token claims",
		})
	}

	return c.Next()
}
