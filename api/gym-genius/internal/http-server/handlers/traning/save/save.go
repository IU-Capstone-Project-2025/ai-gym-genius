package save

import (
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type Request struct {
	Name string `json:"name" validate:"required"`
}

type Response struct {
	Status     string    `json:"status"`
	Error      string    `json:"error,omitempty"`
	TrainingID uuid.UUID `json:"training_id,omitempty"`
}

type TrainingSaver interface {
	Create(name string) (uuid.UUID, error)
}

func New(saver TrainingSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, Response{
				Status: "error",
				Error:  "error decoding request: " + err.Error(),
			})
			return
		}
		if err := validator.New().Struct(req); err != nil {
			render.JSON(w, r, Response{
				Status: "error",
				Error:  "validation error: " + err.Error(),
			})
			return
		}
		id, err := saver.Create(req.Name)
		if err != nil {
			render.JSON(w, r, Response{
				Status: "error",
				Error:  "error creating training: " + err.Error(),
			})
			return
		}
		render.JSON(w, r, Response{
			Status:     "success",
			TrainingID: id,
		})
	}
}
