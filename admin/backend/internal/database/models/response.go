package models

import "time"

type ErrorResponse struct {
	Error string `json:"error" example:"A descriptive error message"`
}

type UserRead struct {
	ID	uint   `json:"id" example:"12345"`
	Login string `json:"username"`
	Name  string `json:"name" example:"John"`
	Surname  string `json:"surname" example:"Doe"`
	Email string `json:"email" example:"john_doe@gmail.com"`
	SubscriptionType string `json:"subscription_type" example:"free"` // e.g., "free", "basic", "pro"
	Status string `json:"status" example:"active"` // e.g., "active", "inactive", "banned"
	LastActivity time.Time `json:"last_activity" example:"2023-10-01T12:00:00Z"`
	NumberOfWorkouts int `json:"number_of_workouts" example:"0"`
	TotalTimeSpent time.Duration `json:"total_time_spent" example:"3600"` // in seconds
	StreakCount int `json:"streak_count" example:"0"`
	AverageWorkoutDuration time.Duration `json:"average_workout_duration" example:"3600"` // in seconds
}

type WorkoutRead struct {
	Duration    time.Duration    `json:"duration" example:"60"`
	Timestamp   time.Time `json:"timestamp" example:"2023-10-01T12:00:00Z"`
	UserID      uint    `json:"user_id" example:"12345"`
}