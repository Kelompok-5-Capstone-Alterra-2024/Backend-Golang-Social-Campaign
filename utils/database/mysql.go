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

	// fundraising := []entities.Fundraising{
	// 	{FundraisingCategoryID: 1,
	// 		OrganizationID:  1,
	// 		Title:           "Bantu Banjir Demak Kembali Pulih",
	// 		ImageUrl:        "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1717434713/207b189d6e16fc148cdb00e012761fda_lscboi.jpg",
	// 		Description:     "Program ini mengajak seluruh lapisan masyarakat untuk berpartisipasi dalam memberikan bantuan berupa kebutuhan pokok, obat-obatan, pakaian, serta bantuan finansial guna mendukung pemulihan dan rehabilitasi wilayah yang terkena dampak. Melalui kolaborasi dengan berbagai organisasi kemanusiaan dan relawan, kegiatan ini diharapkan dapat meringankan beban para korban banjir. Tujuan bantuan ini adalah untuk mempercepat penanganan darurat bencana banjir yang telah merendam sebagian besar wilayah Demak. Banjir di Kabupaten Demak menjadi peringatan bagi pemerintah dan masyarakat untuk meningkatkan kesiapsiagaan dan penanganan bencana.Harapannya kondisi segera membaik dan masyarakat bisa kembali normal dengan bantuan dan kepedulian dari lembaga seperti BIN.",
	// 		Status:          "Started",
	// 		GoalAmount:      15000000,
	// 		CurrentProgress: 10000000,
	// 		StartDate:       time.Now(),
	// 		EndDate:         time.Now().AddDate(0, 0, 20)},
	// }

	organization := &entities.Organization{
		Name:        "Yayasan Berbagai kasih",
		Description: "Tentang Yayasan Berbagi Kasih merupakan Yayasan Independen yang lahir dari inisiatif PeduliPintar untuk mengelola dan mengimplementasikan program sosial secara profesional dan transparan. Sebagai yayasan independen, Social Project tidak hanya mengimplementasikan donasi dari PeduliPintar, melainkan juga berkolaborasi dengan beragam pihak untuk menjalankan, mengimplementasikan, dan menyalurkan donasi beragam program sosial. Telah menghubungkan jutaan dampak sosial. Sampai saat ini, Yayasan Berbagi Kasih telah menjalankan ratusan program sosial dan berkolaborasi dengan ratusan mitra kolaborator. Mulai dari pemulihan bencana alam sampai bantuan kebutuhan pokok. ",
		Avatar:      "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1717434715/d2cb2595f0ac5557f5c6bda8028ce4b4_jixcwa.png",
		IsVerified:  true,
	}

	// category := []entities.FundraisingCategory{
	// 	{Name: "Edukasi"},
	// 	{Name: "Bencana"},
	// 	{Name: "Sosial"},
	// 	{Name: "Alam"},
	// }

	// resultFund := DB.Create(&fundraising)
	// if resultFund.Error != nil {
	// 	panic("failed to create fundraising")
	// }

	resultOrg := DB.Create(&organization)
	if resultOrg.Error != nil {
		panic("failed to create organization")
	}

	// resultCat := DB.Create(&category)
	// if resultCat.Error != nil {
	// 	panic("failed to create category")
	// }

	fmt.Println("Database connection successful!")

}
