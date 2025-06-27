package database

import (
	"admin/internal/database/schemas"
	"admin/config"
	
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var(
	secret = config.C.Secret
	jwtSecret = config.C.JwtSecret
	dbHost = config.C.DbHost
	dbUser = config.C.DbUser
	dbPassword = config.C.DbPassword
	dbName = config.C.DbName
	dbPort = config.C.DbPort
)

func Hash(login, password string) string {
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

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&schemas.User{},
		&schemas.Admin{},
		&schemas.UserActivity{},
		&schemas.Exercise{},
		&schemas.ExerciseSet{},
		&schemas.ExerciseInfo{},
		&schemas.MuscleGroup{},
		&schemas.Workout{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

