package database

import (
	"admin/config"
	"admin/internal/database/schemas"

	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"encoding/json"
	"os"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var (
	secret     = config.C.Secret
	jwtSecret  = config.C.JwtSecret
	dbHost     = config.C.DbHost
	dbUser     = config.C.DbUser
	dbPassword = config.C.DbPassword
	dbName     = config.C.DbName
	dbPort     = config.C.DbPort
)

func Hash(login, password string) string {
	data := login + ":" + password + ":" + secret
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func CreateTokenForUser(user schemas.Admin) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"login": user.Login,
		"role":  "user",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
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

	switch config.C.AppEnv {
	case "PROD":
		DB, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				TranslateError: true,                                  // fix to properly return errors
				Logger:         logger.Default.LogMode(logger.Silent), // silence the gorm logger
			},
		)
		// DB = DB.Debug() // debug postgres queries if needed
	case "DEV":
		DB, err = gorm.Open(
			sqlite.Open("devDb.db"),
			&gorm.Config{
				TranslateError: true, // fix to properly return errors
				// Logger: logger.Default.LogMode(logger.Silent), // silence the gorm logger
			},
		)
		DB = DB.Debug() // outputs generated sql to stdout
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&schemas.User{},
		&schemas.Admin{},
		&schemas.UserActivity{},
		&schemas.Workout{},
		&schemas.Exercise{},
		&schemas.ExerciseSet{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return UpsertStaticData()
}

func UpsertStaticData() error {
	err := UpsertStaticUsers()
	if err != nil {
		return fmt.Errorf("failed to upsert static users: %w", err)
	}

	err = UpsertStaticExercises()
	if err != nil {
		return fmt.Errorf("failed to upsert static exercises: %w", err)
	}

	return nil
}

func UpsertStaticUsers() error {
	data, err := os.ReadFile("assets/users.json")
	if err != nil {
		return err
	}

	var users []schemas.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return err
	}

	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		UpdateAll: true,
	}).Create(&users).Error
	if err != nil {
		return err
	}

	return nil
}

// reads assets/exercises.json and upserts them into the db
func UpsertStaticExercises() error {
	data, err := os.ReadFile("assets/exercises.json")
	if err != nil {
		return err
	}

	var exercises []schemas.Exercise
	err = json.Unmarshal(data, &exercises)
	if err != nil {
		return err
	}

	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	}).Create(&exercises).Error
	if err != nil {
		return err
	}

	return nil
}
