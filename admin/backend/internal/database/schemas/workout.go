package schemas

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	ID           uint          `gorm:"primaryKey"`
	UserID       uint          `gorm:"not null"`
	User         User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Duration     time.Duration `gorm:"not null"`
	StartTime    time.Time     `gorm:"not null"` // Timestamp of the workout
	ExerciseSets []ExerciseSet
}

var (
	ErrUserNotActive = errors.New("user is not active")
)

// check that user is active to create this workout
func (w *Workout) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	if err = tx.First(&user, w.UserID).Error; err != nil {
		return err
	}

	if user.Status != "active" {
		return ErrUserNotActive
	}

	return nil
}

// update user stats
func (w *Workout) AfterCreate(tx *gorm.DB) (err error) {
	result :=
		tx.Model(&User{}).
		Where("id = ?", w.UserID).
		Updates(map[string]any{
			"number_of_workouts":       gorm.Expr("number_of_workouts + 1"),
			"total_time_spent":         gorm.Expr("total_time_spent + ?", w.Duration),
			"average_workout_duration": gorm.Expr("(total_time_spent + ?) / (number_of_workouts + 1)", w.Duration),
			"last_activity":            time.Now(),
		})
	return result.Error
}
