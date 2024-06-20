package dto

import "capstone/entities"

type DistributeFundFundraisingRequest struct {
	FundraisingID uint   `json:"fundraising_id" form:"fundraising_id"`
	Amount        int    `json:"amount" form:"amount"`
	BankName      string `json:"bank_name" form:"bank_name"`
	NoRekening    string `json:"no_rekening" form:"no_rekening"`
	Name          string `json:"name" form:"name"`
	ImagePayment  string `json:"image_payment" form:"image_payment"`
}

type AdminFundraisingsResponse struct {
	ID               uint   `json:"id"`
	Title            string `json:"title"`
	OrganizationName string `json:"organization_name"`
	CurrentProgress  int    `json:"current_progress"`
	TargetAmount     int    `json:"target_amount"`
	Status           string `json:"status"`
}

func ToAdminFundraisingsResponse(fundraising entities.Fundraising) AdminFundraisingsResponse {
	return AdminFundraisingsResponse{
		ID:               fundraising.ID,
		Title:            fundraising.Title,
		OrganizationName: fundraising.Organization.Name,
		CurrentProgress:  fundraising.CurrentProgress,
		TargetAmount:     fundraising.GoalAmount,
		Status:           fundraising.Status,
	}
}

func ToAdminAllFundraisingsResponse(fundraisings []entities.Fundraising) []AdminFundraisingsResponse {
	var result []AdminFundraisingsResponse
	for _, fundraising := range fundraisings {
		result = append(result, ToAdminFundraisingsResponse(fundraising))
	}
	return result
}

type AdminFundraisingResponse struct {
	ID              uint   `json:"id"`
	OrganizationID  uint   `json:"organization_id"`
	CategoryID      uint   `json:"category_id"`
	Title           string `json:"title"`
	CurrentProgress int    `json:"current_progress"`
	TargetAmount    int    `json:"target_amount"`
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
}

func ToAdminFundraisingResponse(fundraising entities.Fundraising) AdminFundraisingResponse {
	return AdminFundraisingResponse{
		ID:              fundraising.ID,
		OrganizationID:  fundraising.OrganizationID,
		CategoryID:      fundraising.FundraisingCategoryID,
		Title:           fundraising.Title,
		CurrentProgress: fundraising.CurrentProgress,
		TargetAmount:    fundraising.GoalAmount,
		ImageUrl:        fundraising.ImageUrl,
		Description:     fundraising.Description,
		StartDate:       fundraising.StartDate.Format("2006-01-02"),
		EndDate:         fundraising.EndDate.Format("2006-01-02"),
	}
}

type AdminFundraisingDonationResponse struct {
	ID                     uint   `json:"id"`
	FundraisingID          uint   `json:"fundraising_id"`
	UserID                 uint   `json:"user_id"`
	UserName               string `json:"user_name"`
	CurrentAmount          int    `json:"current_amount"`
	PaymentMethod          string `json:"payment_method"`
	DonatedDate            string `json:"donated_date"`
	FundraisingDescription string `json:"fundraising_description"`
}

func ToAdminFundraisingDonationResponse(donation entities.DonationManual) AdminFundraisingDonationResponse {
	return AdminFundraisingDonationResponse{
		ID:                     donation.ID,
		FundraisingID:          donation.FundraisingID,
		UserID:                 donation.UserID,
		UserName:               donation.User.Fullname,
		CurrentAmount:          donation.Fundraising.CurrentProgress,
		PaymentMethod:          "Transfer Bank",
		DonatedDate:            donation.CreatedAt.Format("2006-01-02"),
		FundraisingDescription: donation.Fundraising.Description,
	}
}

func ToAdminAllFundraisingDonationResponse(donations []entities.DonationManual) []AdminFundraisingDonationResponse {
	var result []AdminFundraisingDonationResponse
	for _, donation := range donations {
		result = append(result, ToAdminFundraisingDonationResponse(donation))
	}
	return result
}

type AdminOrgResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	IsVerified  bool   `json:"is_verified"`
	JoinDate    string `json:"join_date"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Website     string `json:"website"`
	Instagram   string `json:"instagram"`
	NoRekening  string `json:"no_rek"`
}

func ToAdminOrgResponse(organization entities.Organization) AdminOrgResponse {
	return AdminOrgResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		IsVerified:  organization.IsVerified,
		JoinDate:    organization.CreatedAt.Format("2006-01-02"),
		Description: organization.Description,
		Avatar:      organization.Avatar,
		Website:     organization.Website,
		Instagram:   organization.Instagram,
		NoRekening:  organization.NoRekening,
	}
}

type AdminOrgFundraisingResponse struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	CurrentProgress int    `json:"current_progress"`
	TargetAmount    int    `json:"target_amount"`
	Status          string `json:"status"`
}

func ToAdminOrgFundraisingResponse(fundraising entities.Fundraising) AdminOrgFundraisingResponse {
	return AdminOrgFundraisingResponse{
		ID:              fundraising.ID,
		Title:           fundraising.Title,
		CurrentProgress: fundraising.CurrentProgress,
		TargetAmount:    fundraising.GoalAmount,
		Status:          fundraising.Status,
	}
}

func ToAdminAllOrgFundraisingResponse(fundraisings []entities.Fundraising) []AdminOrgFundraisingResponse {
	var result []AdminOrgFundraisingResponse
	for _, fundraising := range fundraisings {
		result = append(result, ToAdminOrgFundraisingResponse(fundraising))
	}
	return result
}

type AdminOrgVolunteersResponse struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	StartDate           string `json:"start_date"`
	TargetVolunteer     int    `json:"target_volunteer"`
	RegisteredVolunteer int    `json:"registered_volunteer"`
	Status              string `json:"status"`
}

func ToAdminOrgVolunteersResponse(volunteer entities.Volunteer) AdminOrgVolunteersResponse {
	return AdminOrgVolunteersResponse{
		ID:                  volunteer.ID,
		Name:                volunteer.Title,
		StartDate:           volunteer.StartDate.Format("2006-01-02"),
		TargetVolunteer:     volunteer.TargetVolunteer,
		RegisteredVolunteer: volunteer.RegisteredVolunteer,
		Status:              volunteer.Status,
	}
}

func ToAdminAllOrgVolunteersResponse(volunteers []entities.Volunteer) []AdminOrgVolunteersResponse {
	var result []AdminOrgVolunteersResponse
	for _, volunteer := range volunteers {
		result = append(result, ToAdminOrgVolunteersResponse(volunteer))
	}
	return result
}

type AdminOrganizationsResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	JoinDate   string `json:"join_date"`
	Website    string `json:"website"`
	Instagram  string `json:"instagram"`
	NoRekening string `json:"no_rek"`
	IsVerified bool   `json:"is_verified"`
	Avatar     string `json:"avatar"`
}

func ToAdminOrganizationsResponse(organization entities.Organization) AdminOrganizationsResponse {
	return AdminOrganizationsResponse{
		ID:         organization.ID,
		Name:       organization.Name,
		JoinDate:   organization.CreatedAt.Format("2006-01-02"),
		Website:    organization.Website,
		Instagram:  organization.Instagram,
		NoRekening: organization.NoRekening,
		IsVerified: organization.IsVerified,
		Avatar:     organization.Avatar,
	}
}

func ToAdminAllOrganizationsResponse(organizations []entities.Organization) []AdminOrganizationsResponse {
	var result []AdminOrganizationsResponse
	for _, organization := range organizations {
		result = append(result, ToAdminOrganizationsResponse(organization))
	}
	return result
}

type AdminAllUsersResponse struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	NoTelp    string `json:"no_telp"`
	CreatedAt string `json:"created_at"`
}

func ToAdminAllUsersResponse(user entities.User) AdminAllUsersResponse {
	return AdminAllUsersResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		NoTelp:    user.NoTelp,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
	}
}

func ToAdminAllUsersResponses(users []entities.User) []AdminAllUsersResponse {
	var result []AdminAllUsersResponse
	for _, user := range users {
		result = append(result, ToAdminAllUsersResponse(user))
	}
	return result
}

type AdminUserDetailResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	RegisterDate string `json:"register_date"`
	Avatar       string `json:"avatar"`
}

func ToAdminUserDetailResponse(user entities.User) AdminUserDetailResponse {
	return AdminUserDetailResponse{
		ID:           user.ID,
		Username:     user.Username,
		FullName:     user.Fullname,
		Email:        user.Email,
		Phone:        user.NoTelp,
		RegisterDate: user.CreatedAt.Format("2006-01-02"),
		Avatar:       user.Avatar,
	}
}

type AdminUserDonationResponse struct {
	ID               uint   `json:"id"`
	FundraisingName  string `json:"fundraising_name"`
	OrganizationName string `json:"organization_name"`
	Amount           int    `json:"amount"`
	TransactionDate  string `json:"transaction_date"`
}

func ToAdminUserDonationResponse(donation entities.DonationManual) AdminUserDonationResponse {
	return AdminUserDonationResponse{
		ID:               donation.ID,
		FundraisingName:  donation.Fundraising.Title,
		OrganizationName: donation.Fundraising.Organization.Name,
		Amount:           donation.Amount,
		TransactionDate:  donation.CreatedAt.Format("2006-01-02"),
	}
}

func ToAdminAllUserDonationResponse(donations []entities.DonationManual) []AdminUserDonationResponse {
	var result []AdminUserDonationResponse
	for _, donation := range donations {
		result = append(result, ToAdminUserDonationResponse(donation))
	}
	return result
}

type AdminUserVolunteers struct {
	ID               uint   `json:"id"`
	VolunteerName    string `json:"volunteer_name"`
	OrganizationName string `json:"organization_name"`
	RegistrationDate string `json:"registration_date"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
}

func ToAdminUserVolunteers(application entities.Application) AdminUserVolunteers {
	return AdminUserVolunteers{
		ID:               application.VacancyID,
		VolunteerName:    application.Volunteer.Title,
		OrganizationName: application.Volunteer.Organization.Name,
		RegistrationDate: application.CreatedAt.Format("2006-01-02"),
		StartDate:        application.Volunteer.StartDate.Format("2006-01-02"),
		EndDate:          application.Volunteer.EndDate.Format("2006-01-02"),
	}
}

func ToAdminAllUserVolunteers(applications []entities.Application) []AdminUserVolunteers {
	var result []AdminUserVolunteers
	for _, application := range applications {
		result = append(result, ToAdminUserVolunteers(application))
	}
	return result
}

type EditUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	NoTelp   string `json:"no_telp" form:"no_telp"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (req *EditUserRequest) ToEntity() entities.User {
	return entities.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		NoTelp:   req.NoTelp,
		Username: req.Username,
		Password: req.Password,
	}
}

type AdminVolunteersResponse struct {
	ID                  uint   `json:"id"`
	OrganizationName    string `json:"organization_name"`
	Title               string `json:"title"`
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	TargetVolunteer     int    `json:"target_volunteer"`
	RegisteredVolunteer int    `json:"registered_volunteer"`
	Status              string `json:"status"`
}

func ToAdminVolunteersResponse(volunteer entities.Volunteer) AdminVolunteersResponse {
	return AdminVolunteersResponse{
		ID:                  volunteer.ID,
		OrganizationName:    volunteer.Organization.Name,
		Title:               volunteer.Title,
		StartDate:           volunteer.StartDate.Format("2006-01-02"),
		EndDate:             volunteer.EndDate.Format("2006-01-02"),
		TargetVolunteer:     volunteer.TargetVolunteer,
		RegisteredVolunteer: volunteer.RegisteredVolunteer,
		Status:              volunteer.Status,
	}
}

func ToAdminAllVolunteersResponse(volunteers []entities.Volunteer) []AdminVolunteersResponse {
	var result []AdminVolunteersResponse
	for _, volunteer := range volunteers {
		result = append(result, ToAdminVolunteersResponse(volunteer))
	}
	return result
}

type AdminArticleResponses struct {
	ID                uint   `json:"id"`
	Title             string `json:"title"`
	Content           string `json:"content"`
	ImageURL          string `json:"image_url"`
	CreatedAt         string `json:"created_at"`
	TotalComment      int    `json:"total_comment"`
	TotalUserBookmark int    `json:"total_user_bookmark"`
}

func ToAdminArticleResponses(article entities.Article, comment []entities.Comment) AdminArticleResponses {

	totalComment := 0
	for _, c := range comment {
		if c.ArticleID == article.ID {
			totalComment += 1
		}
	}

	return AdminArticleResponses{
		ID:                article.ID,
		Title:             article.Title,
		Content:           article.Content,
		ImageURL:          article.ImageURL,
		CreatedAt:         article.CreatedAt.Format("2006-01-02"),
		TotalComment:      totalComment,
		TotalUserBookmark: 0,
	}
}

func ToAdminAllArticleResponses(articles []entities.Article, comments []entities.Comment) []AdminArticleResponses {
	var result []AdminArticleResponses
	for _, article := range articles {
		result = append(result, ToAdminArticleResponses(article, comments))
	}
	return result
}

type AdminArticleResponse struct {
	ID                uint   `json:"id"`
	Title             string `json:"title"`
	Content           string `json:"content"`
	TotalComment      int    `json:"total_comment"`
	TotalUserBookmark int    `json:"total_user_bookmark"`
}

func ToAdminArticleResponse(article entities.Article, comment []entities.Comment) AdminArticleResponse {
	totalComment := 0

	for _, c := range comment {
		if c.ArticleID == article.ID {
			totalComment += 1
		}
	}

	return AdminArticleResponse{
		ID:                article.ID,
		Title:             article.Title,
		Content:           article.Content,
		TotalComment:      totalComment,
		TotalUserBookmark: 0,
	}
}

type AdminDonationResponses struct {
	ID               uint   `json:"id"`
	UserFullName     string `json:"user_fullname"`
	FundraisingTitle string `json:"fundraising_title"`
	OrganizationName string `json:"organization_name"`
	Amount           int    `json:"amount"`
	ImagePayment     string `json:"image_payment"`
	CreatedAt        string `json:"created_at"`
	Status           string `json:"status"`
}

func ToAdminDonationResponses(donation entities.DonationManual) AdminDonationResponses {
	return AdminDonationResponses{
		ID:               donation.ID,
		UserFullName:     donation.User.Fullname,
		FundraisingTitle: donation.Fundraising.Title,
		OrganizationName: donation.Fundraising.Organization.Name,
		Amount:           donation.Amount,
		ImagePayment:     donation.ImagePayment,
		CreatedAt:        donation.CreatedAt.Format("2006-01-02"),
		Status:           donation.Status,
	}
}

func ToAdminAllDonationResponses(donations []entities.DonationManual) []AdminDonationResponses {
	var result []AdminDonationResponses
	for _, donation := range donations {
		result = append(result, ToAdminDonationResponses(donation))
	}
	return result
}

type TransactionSummary struct {
	Transaction []TransactionData `json:"transaction"`
	TotalAmount float64           `json:"total_amount"`
	Moth        string            `json:"moth"`
	Percentage  float64           `json:"percentage"`
}

type TransactionData struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

func ToTransactionSummary(transactions []entities.Transaction, totalAmount float64, month string, percentage float64) TransactionSummary {
	var transactionData []TransactionData

	for _, transaction := range transactions {
		transactionData = append(transactionData, TransactionData{
			Date:   transaction.CreatedAt.Format("2006-01-02"),
			Amount: float64(transaction.Amount),
		})
	}

	return TransactionSummary{
		Transaction: transactionData,
		TotalAmount: totalAmount,
		Moth:        month,
		Percentage:  percentage,
	}
}
