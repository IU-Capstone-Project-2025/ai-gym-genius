package models

import "time"

type UserCreate struct {
	// Login    string `json:"login"`
	Login string `json:"login" example:"john"`
	// Surname  string `json:"surname"`
	// Phone    string `json:"phone"`
	Password string `json:"password" example:"123"`
}

type WorkoutCreate struct {
	Duration    uint      `json:"duration" example:"60"`
	StartTime   time.Time `json:"start_time" example:"2023-10-01T12:00:00Z"`
	Description string    `json:"description" example:"Morning workout"`
	Weight      float64   `json:"weight" example:"70.5"`
}

type AddExercise struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type UserUpdate struct {
	Login    *string `json:"login"`
	// Surname  *string `json:"surname"`
	// Phone    *string `json:"phone"`
	Password *string `json:"password"`
}

type WorkoutUpdate struct {
	Duration    uint      `json:"duration" example:"60"`
	StartTime   time.Time `json:"start_time" example:"2023-10-01T12:00:00Z"`
	Description string    `json:"description" example:"Morning workout"`
	Weight      float64   `json:"weight" example:"70.5"`
}

type ExerciseUpdate struct {
	Name string `json:"url"`
	URL string `json:"url"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserActivityCreate struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}
