package models

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

type GetNumberOfActiveUsersInput struct {
	StartDate string `json:"start_date" example:"2023-01-01"`
	EndDate   string `json:"end_date" example:"2023-01-31"`
	Step string `json:"step" example:"day"`
}

type GetUserActivityInput struct {
	UserID	string `json:"user_id"`
	Date string `json:"date"`
}