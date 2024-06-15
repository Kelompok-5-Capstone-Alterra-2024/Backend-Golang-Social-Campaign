package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	"capstone/repositories"
	"context"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type DonationService interface {
	CreateDonation(donationRequest dto.DonationRequest) (entities.Donation, error)
	GetDonationByID(id int) (entities.Donation, error)
	GetDonationByUserID(ctx context.Context, limit int, offset int, userID uint) ([]entities.Donation, error)
	GetByFundraisingID(id int) ([]entities.Donation, error)
	GetDonationCommentByFundraisingID(id int) ([]entities.DonationComment, error)
	LikeComment(ctx context.Context, commentID uint, userID uint) error
	UnlikeComment(ctx context.Context, commentID uint, userID uint) error
	PaymentProcess(request dto.TransactionNotificationRequest) error
	FetchStatusFromMidtrans(orderID string) (dto.TransactionNotificationRequest, error)
}

type donationService struct {
	donationRepository repositories.DonationRepository
	fundraisingRepo    repositories.FundraisingRepository
}

func NewDonationService(donationRepository repositories.DonationRepository, fundraisingRepo repositories.FundraisingRepository) *donationService {
	return &donationService{donationRepository, fundraisingRepo}
}

func (s *donationService) CreateDonation(donationRequest dto.DonationRequest) (entities.Donation, error) {
	donation := entities.Donation{
		Amount:        donationRequest.Amount,
		UserID:        donationRequest.User.ID,
		FundraisingID: donationRequest.ID,
		Status:        "pending",
	}

	fundraising, err := s.fundraisingRepo.FindByID(int(donationRequest.ID))
	if err != nil {
		return donation, err
	}

	Random := helper.GenerateRandomOTP(5)
	atoi, err := strconv.Atoi(Random)

	paymentTrans := entities.PaymentTransaction{
		ID:              atoi,
		Amount:          donationRequest.Amount,
		FundraisingName: fundraising.Title,
	}

	paymentUrl, err := helper.GetPaymentUrl(paymentTrans, donationRequest.User)
	if err != nil {
		return donation, err
	}

	donation.PaymentUrl = paymentUrl
	donation.Code = Random

	newDonation, err := s.donationRepository.Save(donation)
	if err != nil {
		return newDonation, err
	}

	comment := entities.DonationComment{
		Comment:    donationRequest.Comment,
		DonationID: newDonation.ID,
		TotalLikes: 0,
	}

	_, err = s.donationRepository.SaveComment(comment)
	if err != nil {
		return newDonation, err
	}

	return newDonation, nil
}

func (s *donationService) GetDonationByUserID(ctx context.Context, limit int, offset int, userID uint) ([]entities.Donation, error) {
	donations, err := s.donationRepository.GetByUserID(limit, offset, userID)
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *donationService) GetDonationByID(id int) (entities.Donation, error) {
	donation, err := s.donationRepository.GetByID(id)
	if err != nil {
		return donation, err
	}
	return donation, nil
}

func (s *donationService) GetDonationCommentByFundraisingID(id int) ([]entities.DonationComment, error) {
	donationComments, err := s.donationRepository.GetCommentsByFundraisingID(id)
	if err != nil {
		return donationComments, err
	}
	return donationComments, nil
}

func (s *donationService) GetByFundraisingID(id int) ([]entities.Donation, error) {
	donations, err := s.donationRepository.GetByFundraisingID(id)
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *donationService) LikeComment(ctx context.Context, commentID uint, userID uint) error {
	liked, err := s.donationRepository.IsLiked(ctx, commentID, userID)
	if err != nil {
		return err
	}
	if liked {
		return nil
	}

	like := entities.LikeDonationComment{
		DonationCommentID: commentID,
		UserID:            userID,
	}

	err = s.donationRepository.CreateLike(ctx, like)
	if err != nil {
		return err
	}

	return s.donationRepository.IncrementLike(ctx, commentID)
}

func (s *donationService) UnlikeComment(ctx context.Context, commentID uint, userID uint) error {

	err := s.donationRepository.RemoveLike(ctx, commentID, userID)
	if err != nil {
		return err
	}

	return s.donationRepository.DecrementLike(ctx, commentID)
}

func (s *donationService) FetchStatusFromMidtrans(orderID string) (dto.TransactionNotificationRequest, error) {
	client := resty.New()
	var statusResponse dto.TransactionNotificationRequest

	url := fmt.Sprintf("https://api.midtrans.com/v2/%s/status", orderID)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Basic U0ItTWlkLXNlcnZlci14X1IzX0JCb0ptU1VfYlJSeGNCV1Y5cGc6").
		SetResult(&statusResponse).
		Get(url)

	if err != nil {
		return statusResponse, err
	}

	if resp.IsError() {
		return statusResponse, fmt.Errorf("error fetching status: %v", resp.String())
	}

	return statusResponse, nil
}

func (s *donationService) PaymentProcess(request dto.TransactionNotificationRequest) error {
	code := request.OrderID

	donation, err := s.donationRepository.GetByCode(code)
	if err != nil {
		return err
	}

	if request.VaNumbers[0].Bank == "bca" {
		donation.PaymentMethod = "BCA Virtual Account"
	} else if request.VaNumbers[0].Bank == "bri" {
		donation.PaymentMethod = "BRI Virtual Account"
	}

	if request.PaymentType == "bank_transfer" && request.TransactionStatus == "capture" && request.FraudStatus == "accept" {
		donation.Status = "paid"
	} else if request.TransactionStatus == "settlement" || request.TransactionStatus == "capture" {
		donation.Status = "paid"
	} else if request.TransactionStatus == "deny" || request.TransactionStatus == "expire" || request.TransactionStatus == "cancel" {
		donation.Status = "failed"
	}

	updatedDonation, err := s.donationRepository.Update(donation)
	if err != nil {
		return err
	}

	fundraising, err := s.fundraisingRepo.FindByID(int(updatedDonation.FundraisingID))
	if err != nil {
		return err
	}

	if updatedDonation.Status == "paid" {
		fundraising.CurrentProgress = fundraising.CurrentProgress + donation.Amount

		if fundraising.CurrentProgress >= fundraising.GoalAmount {
			fundraising.Status = "Achived"
		}

		_, err = s.fundraisingRepo.Update(fundraising)
		if err != nil {
			return err
		}
	}

	return nil

}
