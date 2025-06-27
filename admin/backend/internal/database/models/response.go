package models

type ErrorResponse struct {
	Error string `json:"error" example:"A descriptive error message"`
}

type UserRead struct {
	Login string `json:"username"`
	ID    uint   `json:"id"`
}
