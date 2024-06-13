package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/repositories"
	"context"
)

type DonationManualService interface {
	CreateManualDonation(donationRequest dto.ManualDonationRequest) (entities.DonationManual, error)
	GetDonationManualByID(id int) (entities.DonationManual, error)
	GetDonationManualByUserID(limit int, offset int, userID uint) ([]entities.DonationManual, error)
	LikeComment(ctx context.Context, commentID uint, userID uint) error
	UnlikeComment(ctx context.Context, commentID uint, userID uint) error
	GetByFundraisingID(id int) ([]entities.DonationManual, error)
	GetDonationCommentByFundraisingID(id int) ([]entities.DonationManualComment, error)
}

type donationManualService struct {
	donationManualRepository repositories.DonationManualRepository
	fundraisingRepo          repositories.FundraisingRepository
}

func NewDonationManualService(donationManualRepository repositories.DonationManualRepository, fundraisingRepo repositories.FundraisingRepository) *donationManualService {
	return &donationManualService{donationManualRepository, fundraisingRepo}
}

func (s *donationManualService) CreateManualDonation(donationRequest dto.ManualDonationRequest) (entities.DonationManual, error) {
	donation := entities.DonationManual{
		ImagePayment:  donationRequest.ImagePayment,
		UserID:        donationRequest.User.ID,
		FundraisingID: donationRequest.ID,
		Status:        "pending",
	}
	// fundraising, err := s.fundraisingRepo.FindByID(int(donationRequest.ID))
	// if err != nil {
	// 	return donation, err
	// }

	newDonation, err := s.donationManualRepository.Save(donation)
	if err != nil {
		return newDonation, err
	}

	comment := entities.DonationManualComment{
		Comment:          donationRequest.Comment,
		DonationManualID: newDonation.ID,
		TotalLikes:       0,
	}

	_, err = s.donationManualRepository.SaveComment(comment)
	if err != nil {
		return newDonation, err
	}

	return newDonation, nil
}

func (s *donationManualService) GetDonationManualByID(id int) (entities.DonationManual, error) {
	donation, err := s.donationManualRepository.GetByID(id)
	if err != nil {
		return donation, err
	}

	return donation, nil
}

func (s *donationManualService) GetDonationManualByUserID(limit int, offset int, userID uint) ([]entities.DonationManual, error) {
	donations, err := s.donationManualRepository.GetByUserID(limit, offset, userID)
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *donationManualService) LikeComment(ctx context.Context, commentID uint, userID uint) error {
	liked, err := s.donationManualRepository.IsLiked(ctx, commentID, userID)
	if err != nil {
		return err
	}
	if liked {
		return nil
	}

	like := entities.LikeDonationManualComment{
		DonationManualCommentID: commentID,
		UserID:                  userID,
	}
	err = s.donationManualRepository.CreateLike(ctx, like)
	if err != nil {
		return err
	}

	return s.donationManualRepository.IncrementLike(ctx, commentID)
}

func (s *donationManualService) UnlikeComment(ctx context.Context, commentID uint, userID uint) error {
	err := s.donationManualRepository.RemoveLike(ctx, commentID, userID)
	if err != nil {
		return err
	}
	return s.donationManualRepository.DecrementLike(ctx, commentID)
}

func (s *donationManualService) GetByFundraisingID(id int) ([]entities.DonationManual, error) {
	donations, err := s.donationManualRepository.GetByDonationID(id)
	if err != nil {
		return donations, err
	}
	return donations, nil
}

func (s *donationManualService) GetDonationCommentByFundraisingID(id int) ([]entities.DonationManualComment, error) {
	donationComments, err := s.donationManualRepository.GetCommentsByDonationID(id)
	if err != nil {
		return donationComments, err
	}
	return donationComments, nil
}
