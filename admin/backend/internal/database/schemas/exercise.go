package schemas

// TODO verify schemas

type Exercise struct {
	ID          uint   `gorm:"primaryKey;autoincrement"`
	Name        string `gorm:"not null;uniqueIndex"`
	Description string `gorm:"not null"`
	MuscleGroup string `gorm:"not null"` // e.g., ["chest", "back", "legs"]
	URL         string
}

type ExerciseSet struct {
	Weight     float64  `json:"weight"`
	Reps       uint     `json:"reps"`
	ExerciseID uint     `json:"exercise_id"`
	Exercise   Exercise `json:""`
}
