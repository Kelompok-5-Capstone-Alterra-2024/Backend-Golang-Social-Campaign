package database

import (
	"capstone/entities"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}

	// fundraising := &entities.Fundraising{
	// 	FundraisingCategoryID: 1,
	// 	OrganizationID:        1,
	// 	Title:                 "Judul Fundraising 3",
	// 	ImageUrl:              "http://example.com/image.jpg",
	// 	Description:           "Deskripsi Fundraising",
	// 	Status:                "Started",
	// 	GoalAmount:            1000000,
	// 	CurrentProgress:       500000,
	// 	StartDate:             time.Now(),
	// 	EndDate:               time.Now().AddDate(0, 0, 20),
	// }

	// organization := &entities.Organization{
	// 	Name:        "Yayasan Berbagai kasih",
	// 	Description: "Deskripsi Organisasi",
	// 	Avatar:      "http://example.com/avatar.jpg",
	// 	IsVerified:  true,
	// }

	category := []entities.FundraisingCategory{
		{Name: "Edukasi"},
		{Name: "Bencana"},
		{Name: "Sosial"},
		{Name: "Alam"},
	}

	// resultFund := DB.Create(&fundraising)
	// if resultFund.Error != nil {
	// 	panic("failed to create fundraising")
	// }

	// resultOrg := DB.Create(&organization)
	// if resultOrg.Error != nil {
	// 	panic("failed to create organization")
	// }

	resultCat := DB.Create(&category)
	if resultCat.Error != nil {
		panic("failed to create category")
	}

	fmt.Println("Database connection successful!")

}
