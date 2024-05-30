package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	"capstone/repositories"
)

type DonationService interface {
	CreateDonation(donationRequest dto.DonationRequest) (entities.Donation, error)
}

type donationService struct {
	donationRepository repositories.DonationRepository
}

func NewDonationService(donationRepository repositories.DonationRepository) *donationService {
	return &donationService{donationRepository}
}

func (s *donationService) CreateDonation(donationRequest dto.DonationRequest) (entities.Donation, error) {
	donation := entities.Donation{
		Amount: donationRequest.Amount,
		UserID: donationRequest.User.ID,
		Status: "pending",
	}

	paymentUrl, err := helper.GetPaymentUrl(donation, donation.User)
	if err != nil {
		return donation, err
	}

	donation.PaymentUrl = paymentUrl

	return s.donationRepository.Save(donation)
}
