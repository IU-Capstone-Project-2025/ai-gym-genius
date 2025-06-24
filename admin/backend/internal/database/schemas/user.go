package schemas

type User struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Login   string `gorm:"unique"`
	Hash    string `gorm:"not null"`
}

type Admin struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
}
