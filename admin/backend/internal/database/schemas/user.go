package schemas

// TODO add other fields: email, phone, or whatever

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Login string `gorm:"uniqueIndex;not null"`
	Hash  string `gorm:"not null"`
}

type Admin struct {
	ID    uint   `gorm:"primaryKey"`
	Login string `gorm:"uniqueIndex"`
	Hash  string `gorm:"not null"`
}
