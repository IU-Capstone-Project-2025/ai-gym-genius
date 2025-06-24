package database

import (
	"admin/internal/database/schemas"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var secret = os.Getenv("SECRET")
var jwtSecret = os.Getenv("JWT_SECRET")

func Hash(login, password string) string {

	fmt.Println(secret, jwtSecret)

	data := login + ":" + password + ":" + secret

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}

func CreateTokenForUser(user schemas.Admin) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func InitDatabase() error {
	var err error

	dsn := "host=localhost user=postgres password=PASSWORD dbname=gorm port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&schemas.UserActivity{},
		&schemas.Admin{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

