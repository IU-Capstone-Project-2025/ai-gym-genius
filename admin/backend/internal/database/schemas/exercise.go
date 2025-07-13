package schemas

type Exercise struct {
	ID           uint     `gorm:"primaryKey;autoincrement"`
	Name         string   `gorm:"not null;uniqueIndex"`
	Description  string   `gorm:"not null"`
	MuscleGroups []string `gorm:"type:text;serializer:json;not null"`
	URL          string
}

type ExerciseSet struct {
	Weight     float64 `gorm:"not null"`
	Reps       uint    `gorm:"not null"`
	ExerciseID uint
	Exercise   Exercise
	WorkoutID  uint
}
