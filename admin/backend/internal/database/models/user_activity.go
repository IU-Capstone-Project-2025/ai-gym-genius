package models

type UserActivity struct {
		ID                 uint             `gorm:"primaryKey;autoIncrement"`
		UserID             string           `gorm:"not null;index"`
		Date 			 string           `gorm:"not null;index"`
}