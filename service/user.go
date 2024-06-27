package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/helper"
	middleware "capstone/middlewares"
	"capstone/repositories"
	"errors"
	"fmt"
	"unicode"

	"github.com/alexedwards/argon2id"
)

type UserService interface {
	Register(request dto.RegisterRequest) (entities.User, error)
	Login(request dto.LoginRequest) (entities.User, string, string, error)
	GetUserByID(id uint) (entities.User, error)
	GenerateResetToken(email string) error
	VerifyOTP(otp string) (entities.User, error)
	ResetPassword(otp, newPassword string) error
	GenerateOTP(email string) error
	GetUserProfile(id int) (entities.User, error)
	EditProfile(userid int, request dto.EditProfileRequest) (entities.User, error)
	ChangePassword(userid int, request dto.ChangePasswordRequest) error
	GetHistoryVolunteer(id uint) ([]dto.UserVolunteerHistory, error)
	GetHistoryVolunteerDetail(id int) (dto.UserVolunteerHistoryDetail, error)
	GetHistoryDonation(id uint) ([]dto.UserDonationHistory, error)

	GetUserFundraisingBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkFundraising, error)
	AddUserFundraisingBookmark(id, userId uint) error
	DeleteUserFundraisingBookmark(id uint, userid uint) error

	GetUserArticleBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkArticle, error)
	AddUserArticleBookmark(id, userId uint) error
	DeleteUserArticleBookmark(id uint, userid uint) error

	GetUserVolunteerBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkVolunteerVacancy, error)
	AddUserVolunteerBookmark(id, userId uint) error
	DeleteUserVolunteerBookmark(id uint, userid uint) error

	GetNotificationFundraising() ([]entities.Fundraising, error)
}

type userService struct {
	userRepository repositories.UserRepository
	adminRepo      repositories.AdminRepository
}

func NewUserService(userRepository repositories.UserRepository, adminRepo repositories.AdminRepository) *userService {
	return &userService{userRepository, adminRepo}
}

func (s *userService) Register(request dto.RegisterRequest) (entities.User, error) {

	if request.ConfirmPass != request.Password {
		return entities.User{}, fmt.Errorf("password doesn't match")
	}

	if err := validatePassword(request.Password); err != nil {
		return entities.User{}, err
	}

	userEmail, _ := s.userRepository.FindByEmail(request.Email)
	if userEmail.Email == request.Email {
		return entities.User{}, errors.New("email already exists")
	}

	userName, _ := s.userRepository.FindByUsername(request.Username)
	if userName.Username == request.Username {
		return entities.User{}, errors.New("username already exists")
	}

	userTelp, _ := s.userRepository.FindByNoTelp(request.NoTelp)
	if userTelp.NoTelp == request.NoTelp {
		return entities.User{}, errors.New("phone number already exists")
	}

	userFull, _ := s.userRepository.FindByFullName(request.Fullname)
	if userFull.Fullname == request.Fullname {
		return entities.User{}, errors.New("fullname already exists")
	}

	passwordHash, err := argon2id.CreateHash(request.Password, &argon2id.Params{Memory: 64 * 1024, Iterations: 4, Parallelism: 4, SaltLength: 16, KeyLength: 32})
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: passwordHash,
		NoTelp:   request.NoTelp,
		Avatar:   "https://res.cloudinary.com/dvrhf8d9t/image/upload/v1715517059/default-avatar_yt6eua.png",
	}

	notif := entities.AdminNotification{
		UserName:  user.Username,
		AvatarURL: user.Avatar,
		Message:   fmt.Sprintf("%s Bergabung ke Peduli Pintar", user.Username),
	}

	err = s.adminRepo.CreateNofication(notif)
	if err != nil {
		return user, err
	}

	return s.userRepository.Save(user)
}

func validatePassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen || !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

func (s *userService) Login(request dto.LoginRequest) (entities.User, string, string, error) {
	username := request.Username
	password := request.Password

	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return user, "", "", err
	}

	if user.Password == password {
		accessToken, refreshToken := middleware.GenerateToken(user.ID, user.Username, "user")
		user.Token = accessToken
		return user, accessToken, refreshToken, nil
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.Password)

	if !match {
		return user, "", "", fmt.Errorf("wrong password")
	}

	accessToken, refreshToken := middleware.GenerateToken(user.ID, user.Username, "user")
	user.Token = accessToken

	return user, accessToken, refreshToken, nil
}

func (s *userService) GenerateResetToken(email string) error {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	resetToken := helper.GenerateToken()
	user.ResetToken = resetToken
	err = s.userRepository.Update(user)
	if err != nil {
		return err
	}

	return helper.SendTokenRestPassword(email, resetToken)
}

func (s *userService) GenerateOTP(email string) error {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return errors.New("email not found")
	}
	otp := helper.GenerateRandomOTP(6)
	user.OTP = otp

	err = s.userRepository.Update(user)
	if err != nil {
		return errors.New("email not found")
	}

	return helper.SendOtpResetPassword(email, otp)
}

func (s *userService) VerifyOTP(otp string) (entities.User, error) {
	user, err := s.userRepository.FindByOTP(otp)
	if err != nil {
		return entities.User{}, err
	}

	if user.OTP != otp {
		return entities.User{}, errors.New("invalid or expired OTP")
	}

	return user, nil
}

func (s *userService) ResetPassword(otp, newPassword string) error {
	// user, err := s.userRepository.FindByResetToken(resetToken)
	// if err != nil {
	// 	return err
	// }

	// if user.ResetToken != resetToken {
	// 	return errors.New("invalid or expired reset token")
	// }

	// user.Password = newPassword
	// user.ResetToken = ""
	user, err := s.userRepository.FindByOTP(otp)
	if err != nil {
		return err
	}

	if user.OTP != otp {
		return errors.New("invalid or expired OTP")
	}

	passwordHash, err := argon2id.CreateHash(newPassword, &argon2id.Params{Memory: 64 * 1024, Iterations: 4, Parallelism: 4, SaltLength: 16, KeyLength: 32})
	if err != nil {
		return err
	}

	user.OTP = ""
	user.Password = passwordHash

	return s.userRepository.Update(user)
}

func (s *userService) GetUserByID(id uint) (entities.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) GetUserProfile(id int) (entities.User, error) {
	return s.userRepository.FindByID(uint(id))
}

func (s *userService) EditProfile(userid int, request dto.EditProfileRequest) (entities.User, error) {

	var existingUser entities.User
	existingUser, err := s.userRepository.FindByID(uint(userid))
	if err != nil {
		return entities.User{}, err
	}

	if request.Fullname != "" {
		existingUser.Fullname = request.Fullname
	}
	if request.Username != "" {
		existingUser.Username = request.Username
	}
	if request.Avatar != "" {
		existingUser.Avatar = request.Avatar
	}
	if request.Email != "" {
		existingUser.Email = request.Email
	}

	return existingUser, s.userRepository.UpdateProfile(uint(userid), existingUser)
}

func (s *userService) ChangePassword(userid int, request dto.ChangePasswordRequest) error {
	user, err := s.userRepository.FindByID(uint(userid))
	if err != nil {
		return err
	}

	// if user.Password != request.CurrentPassword {
	// 	return errors.New("wrong current password")
	// }

	// if request.NewPassword != request.ConfirmPassword {
	// 	return errors.New("password doesn't match")
	// }

	// user.Password = request.NewPassword
	// return s.userRepository.Update(user)

	if user.Password == request.CurrentPassword {
		passwordHash, err := argon2id.CreateHash(request.NewPassword, &argon2id.Params{Memory: 64 * 1024, Iterations: 4, Parallelism: 4, SaltLength: 16, KeyLength: 32})
		if err != nil {
			return err
		}

		user.Password = passwordHash
		return s.userRepository.Update(user)
	}

	match, err := argon2id.ComparePasswordAndHash(request.CurrentPassword, user.Password)
	if err != nil {
		return err
	}

	// Jika tidak cocok dengan hash, cek secara langsung
	if !match && user.Password != request.CurrentPassword {
		return errors.New("wrong current password")
	}

	if request.NewPassword != request.ConfirmPassword {
		return errors.New("password doesn't match")
	}

	// Hashing password baru sebelum disimpan
	passwordHash, err := argon2id.CreateHash(request.NewPassword, &argon2id.Params{Memory: 64 * 1024, Iterations: 4, Parallelism: 4, SaltLength: 16, KeyLength: 32})
	if err != nil {
		return err
	}

	user.Password = passwordHash
	return s.userRepository.Update(user)
}

func (s *userService) GetHistoryVolunteer(id uint) ([]dto.UserVolunteerHistory, error) {
	userHistory, err := s.userRepository.GetHistoryVolunteer(id)
	if err != nil {
		return nil, err
	}

	var userVolunteerHistory []dto.UserVolunteerHistory
	for _, history := range userHistory {
		Volunteers, err := s.userRepository.GetVolunteerById(history.VacancyID)
		if err != nil {
			return nil, err
		}
		userVolunteerHistory = append(userVolunteerHistory, dto.UserVolunteerHistory{
			ID:        Volunteers.ID,
			Title:     Volunteers.Title,
			Location:  Volunteers.Location,
			ImageURL:  Volunteers.ImageURL,
			StartDate: Volunteers.StartDate,
			EndDate:   Volunteers.EndDate,
			Status:    Volunteers.Status,
		})
	}

	return userVolunteerHistory, nil
}

func (s *userService) GetHistoryVolunteerDetail(id int) (dto.UserVolunteerHistoryDetail, error) {
	Volunteers, err := s.userRepository.GetVolunteerById(uint(id))
	if err != nil {
		return dto.UserVolunteerHistoryDetail{}, err
	}

	return dto.UserVolunteerHistoryDetail{
		ID:              Volunteers.ID,
		Title:           Volunteers.Title,
		ImageURL:        Volunteers.ImageURL,
		Location:        Volunteers.Location,
		ContentActivity: Volunteers.ContentActivity,
		StartDate:       Volunteers.StartDate,
		EndDate:         Volunteers.EndDate,
		Status:          Volunteers.Status,
	}, nil
}

func (s *userService) GetHistoryDonation(id uint) ([]dto.UserDonationHistory, error) {
	userHistory, err := s.userRepository.GetHistoryDonation(id)
	if err != nil {
		return nil, err
	}

	var userDonationHistory []dto.UserDonationHistory
	for _, history := range userHistory {
		Donations, err := s.userRepository.GetFundraisingById(history.FundraisingID)
		if err != nil {
			return nil, err
		}
		userDonationHistory = append(userDonationHistory, dto.UserDonationHistory{
			ID:       Donations.ID,
			Tittle:   Donations.Title,
			ImageURL: Donations.ImageUrl,
			Status:   Donations.Status,
			Amount:   history.Amount,
		})
	}

	return userDonationHistory, nil
}

func (s *userService) GetUserFundraisingBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkFundraising, error) {
	return s.userRepository.FindUserFundraisingBookmark(id, limit, offset)
}

func (s *userService) AddUserFundraisingBookmark(id, userId uint) error {
	added, err := s.userRepository.IsFundraisingBookmark(id, userId)
	if err != nil {
		return err
	}

	if added {
		return nil
	}

	user := entities.UserBookmarkFundraising{
		UserID:        userId,
		FundraisingID: id,
	}

	return s.userRepository.AddUserFundraisingBookmark(user)
}

func (s *userService) DeleteUserFundraisingBookmark(id uint, userid uint) error {

	return s.userRepository.DeleteUserFundraisingBookmark(id, userid)
}

func (s *userService) GetUserArticleBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkArticle, error) {
	return s.userRepository.FindUserArticleBookmark(id, limit, offset)
}

func (s *userService) AddUserArticleBookmark(id, userId uint) error {
	added, err := s.userRepository.IsArticleBookmark(id, userId)
	if err != nil {
		return err
	}

	if added {
		return nil
	}

	user := entities.UserBookmarkArticle{
		UserID:    userId,
		ArticleID: id,
	}

	return s.userRepository.AddUserArticleBookmark(user)
}

func (s *userService) DeleteUserArticleBookmark(id uint, userid uint) error {

	return s.userRepository.DeleteUserArticleBookmark(id, userid)
}

func (s *userService) GetUserVolunteerBookmark(id uint, limit int, offset int) ([]entities.UserBookmarkVolunteerVacancy, error) {
	return s.userRepository.FindUserVolunteerBookmark(id, limit, offset)
}

func (s *userService) AddUserVolunteerBookmark(id, userId uint) error {
	added, err := s.userRepository.IsVolunteerBookmark(id, userId)
	if err != nil {
		return err
	}

	if added {
		return nil
	}

	user := entities.UserBookmarkVolunteerVacancy{
		UserID:               userId,
		VolunteerVacanciesID: id,
	}

	return s.userRepository.AddUserVolunteerBookmark(user)
}

func (s *userService) DeleteUserVolunteerBookmark(id uint, userid uint) error {

	return s.userRepository.DeleteUserVolunteerBookmark(id, userid)
}

func (s *userService) GetNotificationFundraising() ([]entities.Fundraising, error) {

	return s.userRepository.GetNotificationFundraising()
}
