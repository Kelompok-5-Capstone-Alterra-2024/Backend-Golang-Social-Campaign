package migration

import (
	"capstone/entities"
	"capstone/utils/database"
	"fmt"
	"log"
)

func Migration() {
	// List of all entities to be migrated
	// s
	if err := database.DB.AutoMigrate(&entities.User{}, &entities.Fundraising{}, &entities.FundraisingCategory{}, &entities.Organization{}, &entities.Donation{}, &entities.DonationComment{}, &entities.LikeDonationComment{}, &entities.DonationManual{}, &entities.DonationManualComment{}, &entities.LikeDonationManualComment{}, &entities.Admin{}, &entities.Volunteer{}, &entities.Application{}, &entities.Article{}, &entities.Comment{}, &entities.LikesComment{}, &entities.TestimoniVolunteer{}, &entities.UserBookmarkArticle{}, &entities.UserBookmarkFundraising{}, &entities.UserBookmarkVolunteerVacancy{}, &entities.Transaction{}, &entities.AdminNotification{}, &entities.Chatbot{}); err != nil {
		log.Fatal("Database migration failed")
	}

	// Loop through each entity and migrate
	// for _, entity := range entitiesToMigrate {
	// 	if database.DB.Migrator().HasTable(entity) == false {
	// 		database.DB.Migrator().CreateTable(entity)
	// 	}

	// 	// if err := database.DB.AutoMigrate(entity); err != nil {
	// 	// 	log.Fatalf("Database migration failed for %T: %v", entity, err)
	// 	// }

	// }
	fmt.Println("Successful database migration")
}
