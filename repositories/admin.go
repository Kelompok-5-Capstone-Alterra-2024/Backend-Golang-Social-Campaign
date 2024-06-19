package repositories

import (
	"capstone/entities"
	"time"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (entities.Admin, error)
	FindAllFundraising(limit int, offset int) ([]entities.Fundraising, error)
	Create(fundraising entities.Fundraising) (entities.Fundraising, error)
	DeleteFundraising(id uint) error
	UpdateFundraisingByID(id uint, fundraising entities.Fundraising) (entities.Fundraising, error)
	FindFundraisingByID(id int) (entities.Fundraising, error)
	FindDonationsByFundraisingID(id int, limit int, offset int) ([]entities.DonationManual, error)
	DistributeFundFundraising(id uint, amount int) (entities.Fundraising, error)

	FindOrganizations(limit int, offset int) ([]entities.Organization, error)
	FindOrganizationByID(id int) (entities.Organization, error)
	UpdateOrganizationByID(id uint, organization entities.Organization) (entities.Organization, error)
	DeleteOrganizationByID(id uint) error
	GetFundraisingByOrganizationID(id int, limt, offset int) ([]entities.Fundraising, error)
	GetVolunteerByOrganizationID(id int, limit, offset int) ([]entities.Volunteer, error)

	FindUsers(limit int, offset int) ([]entities.User, error)
	FindUserByID(id int) (entities.User, error)
	UpdateUserByID(id uint, user entities.User) (entities.User, error)
	FindDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error)
	FindVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error)
	DeleteUserWithDonations(id uint) error

	FindAllDonations(page, limit int) ([]entities.DonationManual, int, error)
	AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error)
	FindFundraisingByDonationID(id int) (entities.Fundraising, error)

	// Dashboard
	FindDonationsLastSevenDays() ([]entities.DonationManual, error)
	GetDailyTransactionStats() ([]TransactionStat, error)
	GetTotalAmountDonations() (int, error)
	GetTotalUserVolunteers() (int, error)
	GetTotalArticles() (int, error)
	GetTotalTransactions() (int, error)
	GetArticlesOrderedByBookmarks(limit int) ([]entities.ArticleWithBookmarkCount, error)
	// GetDonationsAmountForPreviousDay() (float64, error)
	// GetVolunteersForPreviousDay() (int64, error)
	// GetArticlesForPreviousDay() (int64, error)
	// GetDonationsForPreviousDay() (int64, error)
	GetTodayDonations() (int64, error)
	GetYesterdayTotalDonations() (int64, error)
	GetTodayVolunteer() (float64, error)
	GetYesterdayTotalVolunteer() (float64, error)
	GetTodayArticle() (float64, error)
	GetYesterdayTotalArticle() (float64, error)
	GetTodayTransaction() (float64, error)
	GetYesterdayTotalTransaction() (float64, error)
	GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByUsername(username string) (entities.Admin, error) {
	var admin entities.Admin
	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *adminRepository) FindAllFundraising(limit int, offset int) ([]entities.Fundraising, error) {
	var fundraisings []entities.Fundraising
	if err := r.db.Preload("Organization").Limit(limit).Offset(offset).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil
}

func (r *adminRepository) Create(fundraising entities.Fundraising) (entities.Fundraising, error) {
	if err := r.db.Create(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) UpdateFundraisingByID(id uint, fundraising entities.Fundraising) (entities.Fundraising, error) {
	if err := r.db.Model(&fundraising).Where("id = ?", id).Updates(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) DeleteFundraising(id uint) error {
	if err := r.db.Delete(&entities.Fundraising{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) FindFundraisingByID(id int) (entities.Fundraising, error) {
	var fundraising entities.Fundraising
	if err := r.db.Preload("Organization").Preload("FundraisingCategory").Where("id = ?", id).First(&fundraising).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) FindDonationsByFundraisingID(id int, limit int, offset int) ([]entities.DonationManual, error) {
	var donations []entities.DonationManual
	if err := r.db.Preload("User").Preload("Fundraising").Where("fundraising_id = ?", id).Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return []entities.DonationManual{}, err
	}
	return donations, nil
}

func (r *adminRepository) DistributeFundFundraising(id uint, amount int) (entities.Fundraising, error) {
	var fundraising entities.Fundraising
	if err := r.db.Model(&fundraising).Where("id = ?", id).Update("current_progress", amount).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return fundraising, nil
}

func (r *adminRepository) FindOrganizations(limit int, offset int) ([]entities.Organization, error) {
	var organizations []entities.Organization
	if err := r.db.Limit(limit).Offset(offset).Order("created_at desc").Find(&organizations).Error; err != nil {
		return []entities.Organization{}, err
	}
	return organizations, nil
}

func (r *adminRepository) FindOrganizationByID(id int) (entities.Organization, error) {
	var organization entities.Organization
	if err := r.db.Where("id = ?", id).First(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *adminRepository) UpdateOrganizationByID(id uint, organization entities.Organization) (entities.Organization, error) {
	if err := r.db.Model(&organization).Where("id = ?", id).Updates(&organization).Error; err != nil {
		return entities.Organization{}, err
	}
	return organization, nil
}

func (r *adminRepository) DeleteOrganizationByID(id uint) error {
	if err := r.db.Delete(&entities.Organization{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetFundraisingByOrganizationID(id int, limit, offest int) ([]entities.Fundraising, error) {

	var fundraisings []entities.Fundraising
	if err := r.db.Preload("Organization").Where("organization_id = ?", id).Offset(offest).Limit(limit).Find(&fundraisings).Error; err != nil {
		return []entities.Fundraising{}, err
	}
	return fundraisings, nil

}

func (r *adminRepository) GetVolunteerByOrganizationID(id int, limit, offest int) ([]entities.Volunteer, error) {
	var volunteers []entities.Volunteer

	if err := r.db.Where("organization_id = ?", id).Offset(offest).Limit(limit).Find(&volunteers).Error; err != nil {
		return []entities.Volunteer{}, err
	}

	return volunteers, nil
}

func (r *adminRepository) FindUsers(limit int, offset int) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return []entities.User{}, err
	}
	return users, nil
}

func (r *adminRepository) FindUserByID(id int) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *adminRepository) UpdateUserByID(id uint, user entities.User) (entities.User, error) {
	if err := r.db.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *adminRepository) FindDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error) {
	var donations []entities.DonationManual
	var total int64
	offset := (page - 1) * limit
	if err := r.db.Preload("Fundraising.Organization").Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&donations).Error; err != nil {
		return []entities.DonationManual{}, 0, err
	}

	r.db.Model(&entities.DonationManual{}).Where("user_id = ?", id).Count(&total)
	return donations, int(total), nil
}

func (r *adminRepository) FindVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error) {

	var applications []entities.Application
	var total int64
	offset := (page - 1) * limit
	if err := r.db.Preload("Volunteer.Organization").Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&applications).Error; err != nil {
		return []entities.Application{}, 0, err
	}

	r.db.Model(&entities.Application{}).Where("user_id = ?", id).Count(&total)
	return applications, int(total), nil

}

func (r *adminRepository) DeleteUserWithDonations(id uint) error {
	if err := r.db.Where("user_id = ?", id).Delete(&entities.Donation{}).Error; err != nil {
		return err
	}

	// Hapus user
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) FindAllDonations(page int, limit int) ([]entities.DonationManual, int, error) {
	var donations []entities.DonationManual
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Preload("User").Preload("Fundraising.Organization").Offset(offset).Limit(limit).Find(&donations).Error; err != nil {
		return []entities.DonationManual{}, 0, err
	}

	r.db.Model(&entities.DonationManual{}).Count(&total)
	return donations, int(total), nil
}

func (r *adminRepository) AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error) {
	var donation entities.DonationManual
	if err := r.db.Model(&donation).Where("id = ?", id).Update("amount", amount).Updates(map[string]interface{}{"status": "sukses"}).Error; err != nil {
		return entities.DonationManual{}, err
	}
	return donation, nil
}

func (r *adminRepository) FindFundraisingByDonationID(id int) (entities.Fundraising, error) {
	// Get Fundraising By Donation ID
	var donation entities.DonationManual
	if err := r.db.Preload("Fundraising").Where("id = ?", id).First(&donation).Error; err != nil {
		return entities.Fundraising{}, err
	}
	return donation.Fundraising, nil
}

func (r *adminRepository) FindDonationsLastSevenDays() ([]entities.DonationManual, error) {
	var donations []entities.DonationManual
	if err := r.db.Where("created_at > ?", time.Now().AddDate(0, 0, -7)).Where("status = ?", "sukses").Find(&donations).Error; err != nil {
		return []entities.DonationManual{}, err
	}
	return donations, nil
}

type TransactionStat struct {
	Date        time.Time `json:"date"`
	TotalAmount float64   `json:"total_amount"`
}

func (r *adminRepository) GetDailyTransactionStats() ([]TransactionStat, error) {
	var stats []TransactionStat

	query := `
        SELECT DATE(created_at) as date, SUM(amount) as total_amount
        FROM transactions
        WHERE created_at >= ?
        GROUP BY DATE(created_at)
        ORDER BY DATE(created_at) ASC
    `

	rows, err := r.db.Raw(query, time.Now().AddDate(0, 0, -7)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var stat TransactionStat
		if err := rows.Scan(&stat.Date, &stat.TotalAmount); err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, nil
}

func (r *adminRepository) GetTotalAmountDonations() (int, error) {
	var total int64
	if err := r.db.Model(&entities.DonationManual{}).Select("SUM(amount)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *adminRepository) GetTotalUserVolunteers() (int, error) {
	var total int64
	if err := r.db.Model(&entities.Application{}).Select("COUNT(id)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *adminRepository) GetTotalArticles() (int, error) {
	var total int64
	if err := r.db.Model(&entities.Article{}).Select("COUNT(id)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *adminRepository) GetTotalTransactions() (int, error) {
	var total int64
	if err := r.db.Model(&entities.Transaction{}).Select("COUNT(id)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}

// func (r *adminRepository) GetDonationsAmountForPreviousDay() (float64, error) {
// 	var total float64
// 	err := r.db.Raw(`
//         SELECT IFNULL(SUM(amount), 0)
//         FROM donation_manuals
//         WHERE DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)
//     `).Scan(&total).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *adminRepository) GetDonationsForPreviousDay() (int64, error) {
// 	var total int64
// 	err := r.db.Raw(`
//         SELECT IFNULL(COUNT(*), 0)
//         FROM transactions
//         WHERE DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)
//     `).Scan(&total).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *adminRepository) GetVolunteersForPreviousDay() (int64, error) {
// 	var total int64
// 	err := r.db.Raw(`
//         SELECT IFNULL(COUNT(*), 0)
//         FROM applications
//         WHERE DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)
//     `).Scan(&total).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

// func (r *adminRepository) GetArticlesForPreviousDay() (int64, error) {
// 	var total int64
// 	err := r.db.Raw(`
//         SELECT IFNULL(COUNT(*), 0)
//         FROM articles
//         WHERE DATE(created_at) = DATE_SUB(CURDATE(), INTERVAL 1 DAY)
//     `).Scan(&total).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }

func (r *adminRepository) GetArticlesOrderedByBookmarks(limit int) ([]entities.ArticleWithBookmarkCount, error) {

	var topArticles []entities.ArticleWithBookmarkCount

	// Query untuk mengambil artikel beserta jumlah bookmarknya
	result := r.db.Table("articles").
		Select("articles.*, COUNT(user_bookmark_articles.article_id) AS bookmark_count").
		Joins("LEFT JOIN user_bookmark_articles ON articles.id = user_bookmark_articles.article_id").
		Group("articles.id").
		Order("bookmark_count DESC").
		Limit(limit).
		Scan(&topArticles)

	if result.Error != nil {
		return nil, result.Error
	}

	return topArticles, nil

}

func (r *adminRepository) GetTodayVolunteer() (float64, error) {
	var todayVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM applications AND DATE(created_at) = CURDATE()
    `, time.Now()).Scan(&todayVolunteers).Error
	return todayVolunteers, err
}

func (r *adminRepository) GetYesterdayTotalVolunteer() (float64, error) {
	var yesterdayTotalVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM applications AND DATE(created_at) < CURDATE()
    `, time.Now()).Scan(&yesterdayTotalVolunteers).Error
	return yesterdayTotalVolunteers, err
}

func (r *adminRepository) GetTodayArticle() (float64, error) {
	var todayVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM articles AND DATE(created_at) = CURDATE()
    `, time.Now()).Scan(&todayVolunteers).Error
	return todayVolunteers, err
}

func (r *adminRepository) GetYesterdayTotalArticle() (float64, error) {
	var yesterdayTotalVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM articles AND DATE(created_at) < CURDATE()
    `, time.Now()).Scan(&yesterdayTotalVolunteers).Error
	return yesterdayTotalVolunteers, err
}

func (r *adminRepository) GetTodayTransaction() (float64, error) {
	var todayVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM transactions AND DATE(created_at) = CURDATE()
    `, time.Now()).Scan(&todayVolunteers).Error
	return todayVolunteers, err
}

func (r *adminRepository) GetYesterdayTotalTransaction() (float64, error) {
	var yesterdayTotalVolunteers float64
	err := r.db.Raw(`
        SELECT COUNT(id) FROM transactions AND DATE(created_at) < CURDATE()
    `, time.Now()).Scan(&yesterdayTotalVolunteers).Error
	return yesterdayTotalVolunteers, err
}

func (r *adminRepository) GetTodayDonations() (int64, error) {
	var todayDonations int64
	err := r.db.Raw(`
        SELECT COALESCE(SUM(amount), 0) 
        FROM donations 
        WHERE status = 'success' 
        AND DATE(created_at) = CURDATE()
    `, time.Now()).Scan(&todayDonations).Error
	return todayDonations, err
}

func (r *adminRepository) GetYesterdayTotalDonations() (int64, error) {
	var yesterdayTotalDonations int64
	err := r.db.Raw(`
        SELECT COALESCE(SUM(amount), 0) 
        FROM donations 
        WHERE status = 'success' 
        AND DATE(created_at) < CURDATE()
    `, time.Now()).Scan(&yesterdayTotalDonations).Error
	return yesterdayTotalDonations, err
}

func (r *adminRepository) GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error) {
	type Result struct {
		ID    uint
		Name  string
		Count int
	}

	var results []Result
	err := r.db.Raw(`
		SELECT fundraising_categories.id, fundraising_categories.name, COUNT(fundraisings.id) as count
		FROM fundraising_categories
		LEFT JOIN fundraisings ON fundraisings.fundraising_category_id = fundraising_categories.id
		GROUP BY fundraising_categories.id
		ORDER BY count DESC
	`).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var categoriesWithCount []entities.FundraisingCategoryWithCount
	for _, result := range results {
		category := entities.FundraisingCategory{}
		category.ID = result.ID
		category.Name = result.Name
		categoriesWithCount = append(categoriesWithCount, entities.FundraisingCategoryWithCount{
			Category: category,
			Count:    result.Count,
		})
	}

	return categoriesWithCount, nil
}
