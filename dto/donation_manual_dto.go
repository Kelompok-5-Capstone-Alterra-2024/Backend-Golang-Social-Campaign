package dto

import "capstone/entities"

type ManualDonationRequest struct {
	ID           uint          `json:"id" `
	ImagePayment string        `json:"image_payment" form:"image_payment"`
	Comment      string        `json:"comment" form:"comment"`
	User         entities.User `json:"user"`
}

type ManualDonationResponse struct {
	ID               uint   `json:"id" `
	UserID           uint   `json:"user_id" `
	CreatedAt        string `json:"created_at" `
	FundraisingName  string `json:"fundraising_name" `
	OrganizationName string `json:"organization_name" `
}

func ToManualDonationResponse(donation entities.DonationManual) ManualDonationResponse {
	return ManualDonationResponse{
		ID:               donation.ID,
		UserID:           donation.UserID,
		CreatedAt:        donation.CreatedAt.String(),
		FundraisingName:  donation.Fundraising.Title,
		OrganizationName: donation.Fundraising.Organization.Name,
	}
}

type HistoryDonationManualResponse struct {
	ID               uint   `json:"id" `
	Amount           int    `json:"amount" `
	FundraisingTitle string `json:"fundraising_title" `
	OrganizationName string `json:"organization_name" `
	Status           string `json:"status" `
	CreatedAt        string `json:"created_at" `
}

func ToHistoryDonationManualResponse(donation entities.DonationManual) HistoryDonationManualResponse {
	return HistoryDonationManualResponse{
		ID:               donation.ID,
		Amount:           donation.Amount,
		FundraisingTitle: donation.Fundraising.Title,
		OrganizationName: donation.Fundraising.Organization.Name,
		Status:           donation.Status,
		CreatedAt:        donation.CreatedAt.Format("2006-01-02"),
	}
}

type DonationManualsResponse struct {
	ID               uint   `json:"id" `
	Amount           int    `json:"amount" `
	ImageUrl         string `json:"image_url" `
	FundraisingTitle string `json:"fundraising_title" `
	Status           string `json:"status" `
	CreatedAt        string `json:"created_at" `
}

func ToDonationManualsResponse(donation entities.DonationManual) DonationManualsResponse {
	return DonationManualsResponse{
		ID:               donation.ID,
		Amount:           donation.Amount,
		ImageUrl:         donation.ImagePayment,
		FundraisingTitle: donation.Fundraising.Title,
		Status:           donation.Status,
		CreatedAt:        donation.CreatedAt.Format("2006-01-02"),
	}
}

func ToDonationManualsResponses(donations []entities.DonationManual) []DonationManualsResponse {
	var response []DonationManualsResponse
	for _, donation := range donations {
		response = append(response, ToDonationManualsResponse(donation))
	}
	return response
}
