package config

import (
	"dumbsound/models"
	"dumbsound/packages/database"
	"fmt"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Artist{},
		&models.Transaction{},
		&models.Music{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
