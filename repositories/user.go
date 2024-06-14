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
