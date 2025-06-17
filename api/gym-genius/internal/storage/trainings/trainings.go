package trainings

import (
	"database/sql"
	"github.com/google/uuid"
)

type TrainingStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *TrainingStorage {
	return &TrainingStorage{db: db}
}

func (s *TrainingStorage) Create(name string) (uuid.UUID, error) {
	return uuid.New(), nil
}
