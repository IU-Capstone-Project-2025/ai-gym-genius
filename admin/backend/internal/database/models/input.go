package models

import "time"

type UserRegistrationInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetUserActivityInput struct {
	UserID uint      `json:"user_id" example:"12345"`
	Date   time.Time `json:"date"`
}
