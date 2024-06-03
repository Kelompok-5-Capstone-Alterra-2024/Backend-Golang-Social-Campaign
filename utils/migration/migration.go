package migration

import (
	"capstone/entities"
	"capstone/utils/database"
	"fmt"
	"log"
)

func Migration() {
	if err := database.DB.AutoMigrate(&entities.User{}, &entities.Fundraising{}, &entities.FundraisingCategory{}, &entities.Organization{}, &entities.Donation{}, &entities.DonationComment{}, &entities.LikeDonationComment{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")
}
