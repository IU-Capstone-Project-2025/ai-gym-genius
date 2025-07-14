package authorization

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"fmt"
)

func AddAdmin(login string, password string) error {
	hash := database.Hash(login, password)

	admin := schemas.Admin{
		Login: login,
		Hash:  hash,
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		panic(err)
	}

	fmt.Println("Admin was successfully added")

	return nil
}
