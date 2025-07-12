package schemas

import (
	"gorm.io/datatypes"
	"time"
)

type Workout struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       uint           `gorm:"not null"`
	User         User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration     time.Duration  `gorm:"not null"`
	Timestamp    time.Time      `gorm:"not null"` // Timestamp of the workout
	ExerciseSets datatypes.JSON `gorm:"type:jsonb;not null"`
}
