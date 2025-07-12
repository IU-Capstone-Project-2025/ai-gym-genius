package models

import "time"

type UserCreate struct {
	Login string `json:"login" example:"john123"`
	Name  string `json:"name" example:"John"`
	Surname  string `json:"surname"`
	Email string `json:"email" example:"john_doe@gmail.com"`
	Password string `json:"password" example:"123"`
}

type WorkoutCreate struct {
	UserID 	uint      `json:"user_id" example:"12345"`
	Duration    time.Duration      `json:"duration" example:"1h45m"` // Duration in hours, minutes, and seconds
	Timestamp time.Time `json:"timestamp" example:"2023-10-01T12:00:00Z"`
}

type ExerciseCreate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ExerciseUpdate struct {
	Name string `json:"url"`
	URL  string `json:"url"`
}

type UserUpdate struct {
	Login *string `json:"login"`
	Name  *string `json:"name" example:"John"`
	Surname  *string `json:"surname"`
	Email *string `json:"email" example:"john_doe@gmail.com"`
	SubscriptionType *string `json:"subscription_type" example:"free"`
	Status *string `json:"status" example:"active"` // e.g., "active", "inactive", "banned"
	LastActivity *time.Time `json:"last_activity" example:"2023-10-01T12:00:00Z"`
	NumberOfWorkouts *int `json:"number_of_workouts" example:"0"`
	TotalTimeSpent *time.Duration `json:"total_time_spent" example:"3600"` // in seconds
	StreakCount *int `json:"streak_count" example:"0"`
	AverageWorkoutDuration *time.Duration `json:"average_workout_duration" example:"3600"` // in seconds
	Password *string `json:"password"`
}

type WorkoutUpdate struct {
	UserID	  *uint      `json:"user_id" example:"12345"`
	Duration    *time.Duration      `json:"duration" example:"60"`
	Timestamp *time.Time `json:"timestamp" example:"2023-10-01T12:00:00Z"`
}


type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserActivityCreate struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}
