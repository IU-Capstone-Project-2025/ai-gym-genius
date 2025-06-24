package schemas

type Admin struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Login   string `gorm:"unique"`
	Hash    string `gorm:"not null"`
}
