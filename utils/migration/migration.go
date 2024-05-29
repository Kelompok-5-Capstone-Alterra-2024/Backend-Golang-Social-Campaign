package migration

import (
	"capstone/entities"
	"capstone/utils/database"
	"fmt"
	"log"
)

func Migration() {
	// List of all entities to be migrated
	entitiesToMigrate := []interface{}{
		&entities.User{},
		&entities.Volunteer{},
		&entities.Application{},
	}

	// Loop through each entity and migrate
	for _, entity := range entitiesToMigrate {
		if err := database.DB.AutoMigrate(entity); err != nil {
			log.Fatalf("Database migration failed for %T: %v", entity, err)
		}
	}

	fmt.Println("Successful database migration")
}
