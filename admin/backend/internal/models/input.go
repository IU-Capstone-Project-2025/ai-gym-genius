package models

import (
	"time"
	"gorm.io/datatypes"
)

type UserCreate struct {
	Login string `json:"login" example:"john123"`
	Name  string `json:"name" example:"John"`
	Surname  string `json:"surname"`
	Email string `json:"email" example:"john_doe@gmail.com"`
	Password string `json:"password" example:"123"`
}

type WorkoutCreate struct {
	UserID 	uint      `json:"user_id" example:"12345"`
	Duration    time.Duration      `json:"duration" swaggerignore:"true" example:"1h45m"` // Duration in hours, minutes, and seconds
	Timestamp time.Time `json:"timestamp" example:"2023-10-01T12:00:00Z"`
    ExerciseSets datatypes.JSON `json:"exercise_sets"`
}

type ExerciseCreate struct {
	Name string `json:"name"`
	Description string `json:"description" example:"Push-ups are a basic exercise that works the chest, shoulders, and triceps."`
	MuscleGroup string `json:"muscle_group" example:"chest"` // e.g., "[chest", "back", "legs]"
	URL  string `json:"url"`
}

type ExerciseUpdate struct {
	Name *string `json:"url"`
	Description *string `json:"description" example:"Push-ups are a basic exercise that works the chest, shoulders, and triceps."`
	MuscleGroup *string `json:"muscle_group" example:"chest"` // e.g., "[chest", "back", "legs]"
	URL *string `json:"image_path" example:"https://example.com/image.jpg"` // URL to the exercise image
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
	TotalTimeSpent *time.Duration `json:"total_time_spent" swaggerignore:"true" example:"3600"` // in seconds
	StreakCount *int `json:"streak_count" example:"0"`
	AverageWorkoutDuration *time.Duration `json:"average_workout_duration" swaggerignore:"true" example:"3600"` // in seconds
	Password *string `json:"password"`
}

type WorkoutUpdate struct {
	UserID	  *uint      `json:"user_id" example:"12345"`
	Duration    *time.Duration      `json:"duration" swaggerignore:"true" example:"60"`
	Timestamp *time.Time `json:"timestamp" example:"2023-10-01T12:00:00Z"`
    ExerciseSets *datatypes.JSON `json:"exercise_sets"`
}


type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ExerciseSetInput struct {
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
	ExerciseID uint   `json:"exercise_id"`
}

type UserActivityCreate struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}

type ExercisesJSON map[string]ExerciseSetInput