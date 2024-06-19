package service

import (
	"capstone/dto"
	"capstone/entities"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"context"
	"fmt"
	"strconv"
	"time"
)

type AdminService interface {
	// Login(request dto.LoginRequest) (entities.Admin, error)
	Login(request dto.LoginRequest) (entities.Admin, string, string, error)
	GetFundraisings(limit int, offset int) ([]entities.Fundraising, error)
	CreateFudraising(ctx context.Context, fundraising entities.Fundraising) (entities.Fundraising, error)
	SaveImageFundraising(id uint, image string) (entities.Fundraising, error)
	DeleteFundraising(id uint) error
	UpdateFundraising(id uint, fundraising entities.Fundraising) (entities.Fundraising, error)
	GetFundraisingByID(id int) (entities.Fundraising, error)
	GetDonationByFundraisingID(id int, limit int, offset int) ([]entities.DonationManual, error)
	DistributeFundFundraising(id uint, amount int) (entities.Fundraising, error)

	GetOrganizations(limit int, offset int) ([]entities.Organization, error)
	GetOrganizationByID(id int) (entities.Organization, error)
	UpdateOrganization(id uint, organization entities.Organization) (entities.Organization, error)
	DeleteOrganization(id uint) error
	GetVolunteerByOrganizationID(id int, limit, offest int) ([]entities.Volunteer, error)
	GetFundraisingByOrganizationID(id int, limit, offest int) ([]entities.Fundraising, error)
	SaveImageOraganization(id uint, image string) (entities.Organization, error)

	GetUsers(limit int, offset int) ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	UpdateUserByID(id uint, user entities.User) (entities.User, error)
	GetDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error)
	GetVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error)
	DeleteUserWithDonations(id uint) error

	GetAllDonations(page, limit int) ([]entities.DonationManual, int, error)
	AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error)

	GetDailyDonationSummary() (map[string]float64, error)
	GetDailyTransactionStats() ([]TransactionData, error)
	GetDataTotalContent() (map[string]interface{}, error)
	GetArticlesOrderedByBookmarks(limit int) ([]entities.ArticleWithBookmarkCount, error)
	GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error)
}

type TransactionData struct {
	Date       time.Time `json:"date"`
	Amount     float64   `json:"amount"`
	Percentage float64   `json:"percentage"`
	Month      string    `json:"month"`
}

type adminService struct {
	adminRepository repositories.AdminRepository
	userRepository  repositories.UserRepository
}

func NewAdminService(adminRepository repositories.AdminRepository, userRepository repositories.UserRepository) *adminService {
	return &adminService{adminRepository, userRepository}
}

func (s *adminService) Login(request dto.LoginRequest) (entities.Admin, string, string, error) {
	username := request.Username
	password := request.Password

	admin, err := s.adminRepository.FindByUsername(username)
	if err != nil {
		return admin, "", "", err
	}

	if admin.Password != password {
		return admin, "", "", fmt.Errorf("wrong password")
	}

	accessToken, refreshToken := middleware.GenerateToken(admin.ID, admin.Username, "admin")
	admin.Token = accessToken

	return admin, accessToken, refreshToken, nil
}

func (s *adminService) GetFundraisings(limit int, offset int) ([]entities.Fundraising, error) {
	return s.adminRepository.FindAllFundraising(limit, offset)
}

func (s *adminService) CreateFudraising(ctx context.Context, fundraising entities.Fundraising) (entities.Fundraising, error) {
	return s.adminRepository.Create(fundraising)
}

func (s *adminService) SaveImageFundraising(id uint, image string) (entities.Fundraising, error) {
	fund, err := s.adminRepository.FindFundraisingByID(int(id))

	if err != nil {
		return fund, err
	}

	fund.ImageUrl = image
	updatedFund, err := s.adminRepository.UpdateFundraisingByID(id, fund)

	if err != nil {
		return updatedFund, err
	}
	return updatedFund, nil
}

func (s *adminService) DeleteFundraising(id uint) error {
	return s.adminRepository.DeleteFundraising(id)
}

func (s *adminService) UpdateFundraising(id uint, fundraising entities.Fundraising) (entities.Fundraising, error) {
	return s.adminRepository.UpdateFundraisingByID(id, fundraising)
}

func (s *adminService) GetFundraisingByID(id int) (entities.Fundraising, error) {
	return s.adminRepository.FindFundraisingByID(id)
}

func (s *adminService) GetDonationByFundraisingID(id int, limit int, offset int) ([]entities.DonationManual, error) {
	return s.adminRepository.FindDonationsByFundraisingID(id, limit, offset)
}

func (s *adminService) DistributeFundFundraising(id uint, amount int) (entities.Fundraising, error) {
	fund, err := s.adminRepository.FindFundraisingByID(int(id))

	if err != nil {
		return fund, err
	}

	donation, err := s.adminRepository.DistributeFundFundraising(id, amount)

	if err != nil {
		return donation, err
	}

	fund.CurrentProgress -= amount
	updatedFund, err := s.adminRepository.UpdateFundraisingByID(id, fund)

	if err != nil {
		return updatedFund, err
	}
	return updatedFund, nil
}

func (s *adminService) GetOrganizations(limit int, offset int) ([]entities.Organization, error) {
	return s.adminRepository.FindOrganizations(limit, offset)
}

func (s *adminService) GetOrganizationByID(id int) (entities.Organization, error) {
	return s.adminRepository.FindOrganizationByID(id)
}

func (s *adminService) UpdateOrganization(id uint, organization entities.Organization) (entities.Organization, error) {
	return s.adminRepository.UpdateOrganizationByID(id, organization)
}

func (s *adminService) DeleteOrganization(id uint) error {
	return s.adminRepository.DeleteOrganizationByID(id)
}

func (s *adminService) GetVolunteerByOrganizationID(id int, limit int, offset int) ([]entities.Volunteer, error) {
	return s.adminRepository.GetVolunteerByOrganizationID(id, limit, offset)
}

func (s *adminService) GetFundraisingByOrganizationID(id int, limit, offest int) ([]entities.Fundraising, error) {
	return s.adminRepository.GetFundraisingByOrganizationID(id, limit, offest)
}

func (s *adminService) SaveImageOraganization(id uint, image string) (entities.Organization, error) {
	org, err := s.adminRepository.FindOrganizationByID(int(id))

	if err != nil {
		return org, err
	}

	org.Avatar = image
	updatedOrg, err := s.adminRepository.UpdateOrganizationByID(id, org)

	if err != nil {
		return updatedOrg, err
	}

	return updatedOrg, nil
}

func (s *adminService) GetUsers(limit int, offset int) ([]entities.User, error) {
	return s.adminRepository.FindUsers(limit, offset)
}

func (s *adminService) GetUserByID(id int) (entities.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *adminService) UpdateUserByID(id uint, user entities.User) (entities.User, error) {
	return s.adminRepository.UpdateUserByID(id, user)
}

func (s *adminService) GetDonationsByUserID(id int, page int, limit int) ([]entities.DonationManual, int, error) {

	return s.adminRepository.FindDonationsByUserID(id, page, limit)

}

func (s *adminService) GetVolunteersByUserID(id int, page int, limit int) ([]entities.Application, int, error) {

	return s.adminRepository.FindVolunteersByUserID(id, page, limit)
}

func (s *adminService) DeleteUserWithDonations(id uint) error {
	return s.adminRepository.DeleteUserWithDonations(id)
}

func (s *adminService) GetAllDonations(page, limt int) ([]entities.DonationManual, int, error) {
	return s.adminRepository.FindAllDonations(page, limt)
}

func (s *adminService) AddAmountToUserDonation(id uint, amount int) (entities.DonationManual, error) {
	fundraising, err := s.adminRepository.FindFundraisingByDonationID(int(id))
	if err != nil {
		return entities.DonationManual{}, err
	}
	donation, err := s.adminRepository.AddAmountToUserDonation(id, amount)

	if err != nil {
		return entities.DonationManual{}, err
	}

	fundraising.CurrentProgress += donation.Amount
	if fundraising.CurrentProgress == fundraising.GoalAmount {
		fundraising.Status = "Achived"
	}
	_, err = s.adminRepository.UpdateFundraisingByID(fundraising.ID, fundraising)
	if err != nil {
		return entities.DonationManual{}, err
	}

	return entities.DonationManual{}, nil

}

func (s *adminService) GetDailyDonationSummary() (map[string]float64, error) {
	donation, err := s.adminRepository.FindDonationsLastSevenDays()

	if err != nil {
		return nil, err
	}

	// data per hari dalam rentang 7 hari terakhir
	dailySummary := make(map[string]float64)
	for _, d := range donation {
		date := d.CreatedAt.Format("2006-01-02")
		dailySummary[date] += float64(d.Amount)
	}

	return dailySummary, nil

}

func (s *adminService) GetDailyTransactionStats() ([]TransactionData, error) {
	stats, err := s.adminRepository.GetDailyTransactionStats()
	if err != nil {
		return nil, err
	}

	var data []TransactionData
	var totalAmountUntilYesterday float64

	for i, stat := range stats {
		percentage := 0.0
		if i > 0 {
			percentage = (stat.TotalAmount / totalAmountUntilYesterday) * 100
		}

		formattedPercentage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", percentage), 64)

		data = append(data, TransactionData{
			Date:       stat.Date,
			Amount:     stat.TotalAmount,
			Percentage: formattedPercentage,
			Month:      stat.Date.Month().String(),
		})
		totalAmountUntilYesterday += stat.TotalAmount
	}

	return data, nil
}

func (s *adminService) GetDataTotalContent() (map[string]interface{}, error) {
	totalAmountDonations, err := s.adminRepository.GetTotalAmountDonations()
	if err != nil {
		return nil, err
	}

	totalUserVolunteers, err := s.adminRepository.GetTotalUserVolunteers()
	if err != nil {
		return nil, err
	}

	totalArticles, err := s.adminRepository.GetTotalArticles()
	if err != nil {
		return nil, err
	}

	totalDonations, err := s.adminRepository.GetTotalTransactions()
	if err != nil {
		return nil, err
	}

	todayDonations, err := s.adminRepository.GetTodayDonations()
	if err != nil {
		return nil, err
	}

	yesterdayDonations, err := s.adminRepository.GetYesterdayTotalDonations()
	if err != nil {
		return nil, err
	}

	todayVolunteer, err := s.adminRepository.GetTodayVolunteer()
	if err != nil {
		return nil, err
	}

	yesterdayVolunteer, err := s.adminRepository.GetYesterdayTotalVolunteer()
	if err != nil {
		return nil, err
	}

	todayArticle, err := s.adminRepository.GetTodayArticle()
	if err != nil {
		return nil, err
	}

	yesterdayArticle, err := s.adminRepository.GetYesterdayTotalArticle()
	if err != nil {
		return nil, err
	}

	todayTransaction, err := s.adminRepository.GetTodayTransaction()
	if err != nil {
		return nil, err
	}

	yesterdayTransaction, err := s.adminRepository.GetYesterdayTotalTransaction()
	if err != nil {
		return nil, err
	}

	percentageDonation := todayDonations / yesterdayDonations * 100
	percentageVolunteer := todayVolunteer / yesterdayVolunteer * 100
	percentageArticle := todayArticle / yesterdayArticle * 100
	percentageTransaction := todayTransaction / yesterdayTransaction * 100

	data := map[string]interface{}{
		"total_donations_amount": totalAmountDonations,
		"persentage_donation":    fmt.Sprintf("%.2f%%", percentageDonation),
		"total_user_volunteers":  totalUserVolunteers,
		"persentage_volunteer":   fmt.Sprintf("%.2f%%", percentageVolunteer),
		"total_articles":         totalArticles,
		"persentage_article":     fmt.Sprintf("%.2f%%", percentageArticle),
		"total_transaction":      totalDonations,
		"persentage_transaction": fmt.Sprintf("%.2f%%", percentageTransaction),
	}

	return data, nil

}

func (s *adminService) GetArticlesOrderedByBookmarks(limit int) ([]entities.ArticleWithBookmarkCount, error) {

	return s.adminRepository.GetArticlesOrderedByBookmarks(limit)
}

// func calculatePercentageChange(current, previous float64) float64 {
// 	if previous == 0 {
// 		if current == 0 {
// 			return 0
// 		}
// 		return 100
// 	}
// 	return ((current - previous) / previous) * 100
// }

func (s *adminService) GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error) {
	return s.adminRepository.GetCategoriesWithCount()
}
