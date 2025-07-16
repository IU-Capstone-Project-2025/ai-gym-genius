package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                     uint          `gorm:"primaryKey"`
	Login                  string        `gorm:"uniqueIndex;not null"`
	Name                   string        `gorm:"not null"`
	Surname                string        `gorm:"not null"`
	Email                  string        `gorm:"not null;uniqueIndex"`
	SubscriptionType       string        `gorm:"not null;check:subscription_type IN ('free', 'basic', 'pro')"`
	Status                 string        `gorm:"not null;check:status IN ('active', 'inactive', 'banned')"`
	LastActivity           time.Time     `gorm:"not null"` // e.g., timestamp of last activity
	NumberOfWorkouts       uint          `gorm:"not null"` // e.g., number of workouts completed
	TotalTimeSpent         time.Duration `gorm:"not null"` // e.g., total time spent in workouts
	StreakCount            uint          `gorm:"not null"` // e.g., number of consecutive workouts
	AverageWorkoutDuration time.Duration `gorm:"not null"` // e.g., average duration of workouts
	Hash                   string        `gorm:"not null" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.SubscriptionType = "free"       // Default subscription type
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
