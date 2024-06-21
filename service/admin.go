package service

import (
	"capstone/dto"
	"capstone/entities"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
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
	GetTransactionsSummary() ([]entities.Transaction, float64, float64, string, error)
	GetDataTotalContent() (map[string]interface{}, error)
	GetArticlesOrderedByBookmarks(limit int) ([]entities.ArticleWithBookmarkCount, error)
	GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error)

	ImportFundraisingFromCSV(reader *csv.Reader) error
	ImportFundraisingFromExcel(file *excelize.File) error

	GetNotificationForAdmin() ([]entities.AdminNotification, error)
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

func (s *adminService) GetTransactionsSummary() ([]entities.Transaction, float64, float64, string, error) {
	transactions, err := s.adminRepository.GetTransactionsLast7Days()
	if err != nil {
		return nil, 0, 0, "", err
	}

	var totalAmount float64
	for _, transaction := range transactions {
		totalAmount += float64(transaction.Amount)
	}

	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)

	totalAmountToday, err := s.adminRepository.GetTotalAmountByDate(today)
	if err != nil {
		return nil, 0, 0, "", err
	}

	totalAmountYesterday, err := s.adminRepository.GetTotalAmountByDate(yesterday)
	if err != nil {
		return nil, 0, 0, "", err
	}

	percentage := 0.0
	if totalAmountYesterday > 0 {
		percentage = (totalAmountToday / totalAmountYesterday) * 100
	}

	month := today.Month().String()

	return transactions, totalAmount, percentage, month, nil
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

func (s *adminService) GetCategoriesWithCount() ([]entities.FundraisingCategoryWithCount, error) {
	return s.adminRepository.GetCategoriesWithCount()
}

func (s *adminService) ImportFundraisingFromCSV(reader *csv.Reader) error {
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		goalAmount, _ := strconv.Atoi(record[4])
		startDate, _ := time.Parse("2006-01-02", record[5])
		endDate, _ := time.Parse("2006-01-02", record[6])

		fundraisingCategoryID, _ := strconv.ParseUint(record[7], 10, 64)
		fundraisingOrgID, _ := strconv.ParseUint(record[8], 10, 64)

		fundraising := entities.Fundraising{
			Title:                 record[1],
			ImageUrl:              record[2],
			Description:           record[3],
			GoalAmount:            goalAmount,
			StartDate:             &startDate,
			EndDate:               &endDate,
			Status:                "unachieved",
			FundraisingCategoryID: uint(fundraisingCategoryID),
			OrganizationID:        uint(fundraisingOrgID),
		}
		s.adminRepository.Create(fundraising)
	}

	return nil
}

func (s *adminService) ImportFundraisingFromExcel(file *excelize.File) error {
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for _, row := range rows[1:] { // Skip header row
		goalAmount, _ := strconv.Atoi(row[4])
		startDate, _ := time.Parse("2006-01-02", row[6])
		endDate, _ := time.Parse("2006-01-02", row[7])

		fundraising := entities.Fundraising{
			Title:       row[1],
			ImageUrl:    row[2],
			Description: row[3],
			GoalAmount:  goalAmount,
			StartDate:   &startDate,
			EndDate:     &endDate,
		}
		s.adminRepository.Create(fundraising)
	}

	return nil
}

func (s *adminService) GetNotificationForAdmin() ([]entities.AdminNotification, error) {

	return s.adminRepository.FindNotifications()
}
