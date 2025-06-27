package models

import "time"

type UserCreate struct {
	// Login    string `json:"login"`
	Login string `json:"login"`
	// Surname  string `json:"surname"`
	// Phone    string `json:"phone"`
	Password string `json:"password"`
}

// includes all fields of UserCreate but optional
type UserUpdate struct {
	Login *string `json:"login"`
	// Surname  *string `json:"surname"`
	// Phone    *string `json:"phone"`
	Password *string `json:"password"`
	ID       uint    `json:"id"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserActivityCreate struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}
