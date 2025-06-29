package models

import "time"

type ErrorResponse struct {
	Error string `json:"error" example:"A descriptive error message"`
}

type UserRead struct {
	Login string `json:"username"`
	ID    uint   `json:"id"`
}

type WorkoutRead struct {
	Duration    uint    `json:"duration" example:"60"`
	StartTime   time.Time  `json:"start_time" example:"2023-10-01T12:00:00Z"`
	Description string  `json:"description" example:"Morning workout"`
	Weight      float64 `json:"weight" example:"70.5"`
}
