package migration

import (
	"capstone/entities"
	"capstone/utils/database"
	"fmt"
)

func Migration() {
	// List of all entities to be migrated
	entitiesToMigrate := []interface{}{
		&entities.User{},
		&entities.Volunteer{},
		&entities.Application{},
		&entities.Article{},
		&entities.Comment{},
		&entities.LikesComment{},
		&entities.TestimoniVolunteer{},
		&entities.Fundraising{},
		&entities.FundraisingCategory{},
		&entities.Organization{},
		&entities.Donation{},
		&entities.DonationComment{},
		&entities.LikeDonationComment{},
	}

	// Loop through each entity and migrate
	for _, entity := range entitiesToMigrate {
		if database.DB.Migrator().HasTable(entity) == false {
			database.DB.Migrator().CreateTable(entity)
		}

		// if err := database.DB.AutoMigrate(entity); err != nil {
		// 	log.Fatalf("Database migration failed for %T: %v", entity, err)
		// }
		// if err := database.DB.AutoMigrate(&entities.User{}, &entities.Fundraising{}, &entities.FundraisingCategory{}, &entities.Organization{}, &entities.Donation{}, &entities.DonationComment{}, &entities.LikeDonationComment{}); err != nil {
		// 	log.Fatal("Database migration failed")
		// }

		fmt.Println("Successful database migration")
	}
}
