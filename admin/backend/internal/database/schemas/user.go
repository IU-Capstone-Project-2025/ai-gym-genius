package schemas

import (
	"admin/config"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var (
	UserStatusExists         = map[string]bool{"active": true, "inactive": true, "banned": true}
	SubscriptionStatusExists = map[string]bool{"active": true, "expired": true, "cancelled": true}
	SubscriptionPlanExists   = map[string]bool{"free": true, "basic": true, "pro": true}
)

type User struct {
	ID                     uint          `gorm:"primaryKey"`
	Login                  string        `gorm:"uniqueIndex;not null"`
	Name                   string        `gorm:"not null"`
	Surname                string        `gorm:"not null"`
	Email                  string        `gorm:"not null;uniqueIndex"`
	SubscriptionPlan       string        `gorm:"not null;check:subscription_plan IN ('free', 'basic', 'pro')"`                        // free, basic, pro
	SubscriptionStatus     string        `gorm:"not null;check:subscription_status IN ('active', 'expired', 'cancelled', 'pending')"` //active, expired, pending
	Status                 string        `gorm:"not null;check:status IN ('active', 'inactive', 'banned')"`                           // active, inactive, pending
	LastActivity           time.Time     `gorm:"not null"`                                                                            // e.g., timestamp of last activity
	NumberOfWorkouts       uint          `gorm:"not null"`                                                                            // e.g., number of workouts completed
	TotalTimeSpent         time.Duration `gorm:"not null"`                                                                            // e.g., total time spent in workouts
	StreakCount            uint          `gorm:"not null"`                                                                            // e.g., number of consecutive workouts
	AverageWorkoutDuration time.Duration `gorm:"not null"`                                                                            // e.g., average duration of workouts
	Hash                   string        `gorm:"not null" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.SubscriptionPlan = "basic"       // Default subscription type
	u.SubscriptionStatus = "active" // Default subscription status
	u.SubscriptionStatus = "active"
	u.Status = "active"               // Default status
	u.LastActivity = time.Now().UTC() // Set current time as last activity
	u.NumberOfWorkouts = 0            // Initial number of workouts
	u.TotalTimeSpent = 0              // Initial total time spent
	u.StreakCount = 0                 // Initial streak count
	u.AverageWorkoutDuration = 0      // Initial average workout duration
	return nil
}

type Admin struct {
	ID    uint   `gorm:"primaryKey"`
	Login string `gorm:"uniqueIndex"`
	Hash  string `gorm:"not null"`
}


func Hash(login, password string) string {
	data := login + ":" + password + ":" + config.C.Secret
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (a *Admin) CreateToken() (string, error) {
	claims := jwt.MapClaims{
		"id":    a.ID,
		"login": a.Login,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.C.JwtSecret))
}

func (u *User) CreateToken() (string, error) {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"login": u.Login,
		"role":  "user",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.C.JwtSecret))
}