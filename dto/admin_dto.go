package dto

import "capstone/entities"

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
	Title           string `json:"title"`
	CurrentProgress int    `json:"current_progress"`
	TargetAmount    int    `json:"target_amount"`
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	EndDate         string `json:"end_date"`
}

func ToAdminFundraisingResponse(fundraising entities.Fundraising) AdminFundraisingResponse {
	return AdminFundraisingResponse{
		ID:              fundraising.ID,
		Title:           fundraising.Title,
		CurrentProgress: fundraising.CurrentProgress,
		TargetAmount:    fundraising.GoalAmount,
		ImageUrl:        fundraising.ImageUrl,
		Description:     fundraising.Description,
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

func ToAdminFundraisingDonationResponse(donation entities.Donation) AdminFundraisingDonationResponse {
	return AdminFundraisingDonationResponse{
		ID:                     donation.ID,
		FundraisingID:          donation.FundraisingID,
		UserID:                 donation.UserID,
		UserName:               donation.User.Fullname,
		CurrentAmount:          donation.Fundraising.CurrentProgress,
		PaymentMethod:          donation.PaymentMethod,
		DonatedDate:            donation.CreatedAt.Format("2006-01-02"),
		FundraisingDescription: donation.Fundraising.Description,
	}
}

func ToAdminAllFundraisingDonationResponse(donations []entities.Donation) []AdminFundraisingDonationResponse {
	var result []AdminFundraisingDonationResponse
	for _, donation := range donations {
		result = append(result, ToAdminFundraisingDonationResponse(donation))
	}
	return result
}

type AdminOrganizationsResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	JoinDate   string `json:"join_date"`
	Contact    string `json:"contact"`
	IsVerified bool   `json:"is_verified"`
}

func ToAdminOrganizationsResponse(organization entities.Organization) AdminOrganizationsResponse {
	return AdminOrganizationsResponse{
		ID:         organization.ID,
		Name:       organization.Name,
		JoinDate:   organization.CreatedAt.Format("2006-01-02"),
		Contact:    organization.Contact,
		IsVerified: organization.IsVerified,
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
	ID           uint                        `json:"id"`
	Username     string                      `json:"username"`
	Email        string                      `json:"email"`
	Phone        string                      `json:"phone"`
	RegisterDate string                      `json:"register_date"`
	Avatar       string                      `json:"avatar"`
	Donations    []AdminUserDonationResponse `json:"donations"`
}

type AdminUserDonationResponse struct {
	DonationID       uint   `json:"donation_id"`
	FundraisingID    uint   `json:"fundraising_id"`
	Title            string `json:"title"`
	OrganizationName string `json:"organization_name"`
	Amount           int    `json:"amount"`
	TransactionDate  string `json:"transaction_date"`
}
