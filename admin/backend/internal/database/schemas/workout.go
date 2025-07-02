package schemas

import "time"

type Workout struct {
	ID          uint      `gorm:"primaryKey"`
	Duration    uint      `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	Description string
	Weight      float64
}
