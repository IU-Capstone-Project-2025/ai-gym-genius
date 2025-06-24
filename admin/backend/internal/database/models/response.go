package models

type ErrorResponse struct {
    Error string `json:"error" example:"A descriptive error message"`
}