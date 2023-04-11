package database

import (
	"fmt"
	"project/models"
)

func RunMigration() {
	err := ConnDB.AutoMigrate(
		&models.Member{},
		&models.Product{},
		&models.ReviewProduct{},
		&models.LikeReview{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
