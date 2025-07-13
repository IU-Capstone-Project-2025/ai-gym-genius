package schemas

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	ID           uint          `gorm:"primaryKey"`
	UserID       uint          `gorm:"not null"`
	User 		 User          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Duration     time.Duration `gorm:"not null"`
	StartTime    time.Time     `gorm:"not null"`
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
    var user User
    if err := tx.First(&user, w.UserID).Error; err != nil {
        return err
    }
	user.NumberOfWorkouts++
	user.TotalTimeSpent += w.Duration
	user.AverageWorkoutDuration = user.TotalTimeSpent / time.Duration(user.NumberOfWorkouts)

    return tx.Save(user).Error
}