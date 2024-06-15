package dto

import (
	"capstone/entities"
)

type VaNumber struct {
	VaNumber string `json:"va_number"`
	Bank     string `json:"bank"`
}

type TransactionNotificationRequest struct {
	FraudStatus       string     `json:"fraud_status" `
	TransactionStatus string     `json:"transaction_status" `
	PaymentType       string     `json:"payment_type" `
	OrderID           string     `json:"order_id" `
	VaNumbers         []VaNumber `json:"va_numbers"`
}

type DonationRequest struct {
	ID      uint          `json:"id" `
	Amount  int           `json:"amount" form:"amount"`
	Comment string        `json:"comment" form:"comment"`
	User    entities.User `json:"user"`
}

type DonationResponse struct {
	ID            uint   `json:"id" `
	Amount        int    `json:"amount" `
	UserID        uint   `json:"user_id" `
	FundraisingID uint   `json:"fundraising_id" `
	Status        string `json:"status" `
	PaymentUrl    string `json:"payment_url" `
}

func ToDonationResponse(donation entities.Donation) DonationResponse {
	return DonationResponse{
		ID:            donation.ID,
		Amount:        donation.Amount,
		UserID:        donation.UserID,
		FundraisingID: donation.FundraisingID,
		Status:        donation.Status,
		PaymentUrl:    donation.PaymentUrl,
	}
}

type HistoryDonationResponse struct {
	ID               uint   `json:"id" `
	Amount           int    `json:"amount" `
	FundraisingTitle string `json:"fundraising_title" `
	OrganizationName string `json:"organization_name" `
	Status           string `json:"status" `
	PaymentMethod    string `json:"payment_method" `
	DonationID       string `json:"donation_id" `
	CreatedAt        string `json:"created_at" `
}

func ToHistoryDonationResponse(donation entities.Donation) HistoryDonationResponse {
	return HistoryDonationResponse{
		ID:               donation.ID,
		Amount:           donation.Amount,
		FundraisingTitle: donation.Fundraising.Title,
		OrganizationName: donation.Fundraising.Organization.Name,
		Status:           donation.Status,
		PaymentMethod:    donation.PaymentMethod,
		DonationID:       donation.Code,
		CreatedAt:        donation.CreatedAt.Format("2006-01-02"),
	}
}

type DonationsResponse struct {
	ID               uint   `json:"id" `
	Amount           int    `json:"amount" `
	ImageUrl         string `json:"image_url" `
	FundraisingTitle string `json:"fundraising_title" `
	Status           string `json:"status" `
	CreatedAt        string `json:"created_at" `
}

func ToDonationsResponse(donations entities.Donation) DonationsResponse {
	return DonationsResponse{
		ID:               donations.ID,
		Amount:           donations.Amount,
		ImageUrl:         donations.Fundraising.ImageUrl,
		FundraisingTitle: donations.Fundraising.Title,
		Status:           donations.Status,
		CreatedAt:        donations.CreatedAt.Format("2006-01-02"),
	}
}

func ToAllDonationsResponses(donations []entities.Donation) []DonationsResponse {
	var result []DonationsResponse
	for _, donation := range donations {
		result = append(result, ToDonationsResponse(donation))
	}
	return result
}
