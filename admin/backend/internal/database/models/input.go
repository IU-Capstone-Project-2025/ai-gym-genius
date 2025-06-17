package models

type GetNumberOfActiveUsersInput struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Step string `json:"step"`
}

type GetUserActivityInput struct {
	UserID	string `json:"user_id"`
	Date string `json:"date"`
}