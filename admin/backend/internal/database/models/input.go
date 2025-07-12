package models

import "time"

type UserCreate struct {
	Login string `json:"login" example:"john123"`
	Name  string `json:"name" example:"John"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone" example:"+1234567890"`
	Email string `json:"email" example:"john_doe@gmail.com"`
	Password string `json:"password" example:"123"`
}

type WorkoutCreate struct {
	Duration    uint      `json:"duration" example:"60"`
	StartTime   time.Time `json:"start_time" example:"2023-10-01T12:00:00Z"`
	Description string    `json:"description" example:"Morning workout"`
	Weight      float64   `json:"weight" example:"70.5"`
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


type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserActivityCreate struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}
