package schemas

import "time"

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Login string `gorm:"uniqueIndex;not null"`
	Name  string `gorm:"not null"`
	Surname string `gorm:"not null"`
	Email string `gorm:"not null;uniqueIndex"`
	SubscriptionType string `gorm:"not null"` // e.g., "free", "basic", "pro"
	Status string `gorm:"not null"` // e.g., "active", "inactive", "banned"
	LastActivity time.Time `gorm:"not null"` // e.g., timestamp of last activity
	NumberOfWorkouts int `gorm:"not null"` // e.g., number of workouts completed
	TotalTimeSpent time.Duration `gorm:"not null"` // e.g., total time spent in workouts
	StreakCount int `gorm:"not null"` // e.g., number of consecutive workouts
	AverageWorkoutDuration time.Duration `gorm:"not null"` // e.g., average duration of workouts
	Hash  string `gorm:"not null"`
}

type Admin struct {
	ID    uint   `gorm:"primaryKey"`
	Login string `gorm:"uniqueIndex"`
	Hash  string `gorm:"not null"`
}
