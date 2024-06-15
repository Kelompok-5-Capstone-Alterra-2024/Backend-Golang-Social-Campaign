package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entities.User) (entities.User, error)
	FindByUsername(username string) (entities.User, error)
	FindByID(id uint) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByFullName(fullname string) (entities.User, error)
	FindByNoTelp(notel string) (entities.User, error)
	FindByResetToken(token string) (entities.User, error)
	FindByOTP(otp string) (entities.User, error)
	Update(user entities.User) error
	UpdateProfile(userid uint, user entities.User) error
	GetHistoryVolunteer(id uint) ([]entities.Application, error)
	GetVolunteerById(id uint) (entities.Volunteer, error)
	GetHistoryDonation(id uint) ([]entities.DonationManual, error)
	GetFundraisingById(id uint) (entities.Fundraising, error)

	FindUserFundraisingBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkFundraising, error)
	AddUserFundraisingBookmark(user entities.UserBookmarkFundraising) error
	DeleteUserFundraisingBookmark(id uint, userid uint) error
	IsFundraisingBookmark(id uint, userid uint) (bool, error)

	FindUserArticleBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkArticle, error)
	AddUserArticleBookmark(user entities.UserBookmarkArticle) error
	DeleteUserArticleBookmark(id uint, userid uint) error
	IsArticleBookmark(id uint, userid uint) (bool, error)

	FindUserVolunteerBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkVolunteerVacancy, error)
	AddUserVolunteerBookmark(user entities.UserBookmarkVolunteerVacancy) error
	DeleteUserVolunteerBookmark(id uint, userid uint) error
	IsVolunteerBookmark(id uint, userid uint) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entities.User) (entities.User, error) {
	if err := r.db.Omit("ResetTokenExpire").Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(username string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByFullName(fullname string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("fullname = ?", fullname).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByNoTelp(notel string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("no_telp = ?", notel).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uint) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByResetToken(token string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("reset_token = ?", token).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByOTP(otp string) (entities.User, error) {
	var user entities.User
	if err := r.db.Where("otp = ?", otp).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user entities.User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) UpdateProfile(userid uint, user entities.User) error {
	return r.db.Model(&entities.User{}).Where("id = ?", userid).Updates(user).Error
}

func (r *userRepository) UpdatePassword(id uint, password string) error {
	return r.db.Model(&entities.User{}).Where("id = ?", id).Update("password", password).Error
}

func (r *userRepository) GetHistoryVolunteer(id uint) ([]entities.Application, error) {
	var application []entities.Application
	if err := r.db.Where("user_id = ?", id).Preload("Volunteer").Find(&application).Error; err != nil {
		return application, err
	}
	return application, nil
}

func (r *userRepository) GetVolunteerById(id uint) (entities.Volunteer, error) {
	var volunteer entities.Volunteer
	if err := r.db.Where("id = ?", id).First(&volunteer).Error; err != nil {
		return volunteer, err
	}
	return volunteer, nil
}

func (r *userRepository) GetHistoryDonation(id uint) ([]entities.DonationManual, error) {
	var donation []entities.DonationManual
	if err := r.db.Where("user_id = ?", id).Preload("Fundraising").Find(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *userRepository) GetFundraisingById(id uint) (entities.Fundraising, error) {
	var donation entities.Fundraising
	if err := r.db.Where("id = ?", id).First(&donation).Error; err != nil {
		return donation, err
	}
	return donation, nil
}

func (r *userRepository) FindUserFundraisingBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkFundraising, error) {
	var bookmark []entities.UserBookmarkFundraising

	if err := r.db.Preload("Fundraising.FundraisingCategory").Limit(limit).Offset(offset).Where("user_id = ?", id).Find(&bookmark).Error; err != nil {
		return bookmark, err
	}
	return bookmark, nil
}

func (r *userRepository) AddUserFundraisingBookmark(user entities.UserBookmarkFundraising) error {

	return r.db.Create(&user).Error
}

func (r *userRepository) DeleteUserFundraisingBookmark(id uint, userid uint) error {

	return r.db.Where("fundraising_id = ? AND user_id = ?", id, userid).Delete(&entities.UserBookmarkFundraising{}).Error
}

func (r *userRepository) IsFundraisingBookmark(id uint, userid uint) (bool, error) {

	var count int64
	if err := r.db.Model(&entities.UserBookmarkFundraising{}).Where("fundraising_id = ? AND user_id = ?", id, userid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) FindUserArticleBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkArticle, error) {

	var bookmark []entities.UserBookmarkArticle

	if err := r.db.Preload("Article").Limit(limit).Offset(offset).Where("user_id = ?", id).Find(&bookmark).Error; err != nil {
		return bookmark, err
	}
	return bookmark, nil
}

func (r *userRepository) AddUserArticleBookmark(user entities.UserBookmarkArticle) error {

	return r.db.Create(&user).Error
}

func (r *userRepository) DeleteUserArticleBookmark(id uint, userid uint) error {

	return r.db.Where("article_id = ? AND user_id = ?", id, userid).Delete(&entities.UserBookmarkArticle{}).Error
}

func (r *userRepository) IsArticleBookmark(id uint, userid uint) (bool, error) {

	var count int64

	if err := r.db.Model(&entities.UserBookmarkArticle{}).Where("article_id = ? AND user_id = ?", id, userid).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) FindUserVolunteerBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkVolunteerVacancy, error) {

	var bookmark []entities.UserBookmarkVolunteerVacancy

	if err := r.db.Preload("Volunteer.Oraganization").Limit(limit).Offset(offset).Where("user_id = ?", id).Find(&bookmark).Error; err != nil {
		return bookmark, err
	}

	return bookmark, nil
}

func (r *userRepository) AddUserVolunteerBookmark(user entities.UserBookmarkVolunteerVacancy) error {

	return r.db.Create(&user).Error
}

func (r *userRepository) DeleteUserVolunteerBookmark(id uint, userid uint) error {

	return r.db.Where("volunteer_vacancies_id = ? AND user_id = ?", id, userid).Delete(&entities.UserBookmarkVolunteerVacancy{}).Error

}

func (r *userRepository) IsVolunteerBookmark(id uint, userid uint) (bool, error) {

	var count int64

	if err := r.db.Model(&entities.UserBookmarkVolunteerVacancy{}).Where("volunteer_vacancies_id = ? AND user_id = ?", id, userid).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
