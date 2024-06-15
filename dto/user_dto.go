package dto

import (
	"capstone/entities"
	"time"
)

type RegisterRequest struct {
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NoTelp      string `json:"no_telp"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	OTP         string `json:"otp"`
	Password    string `json:"new_password"`
	ConfirmPass string `json:"confirm_password"`
}

type UserProfileResponse struct {
	ID       uint   `json:"user_id"`
	Avatar   string `json:"avatar_url"`
	Username string `json:"username"`
}

type EditProfileRequest struct {
	// ID       uint   `json:"user_id"`
	Fullname string `json:"full_name" form:"full_name"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar_url" form:"avatar_url"`
	NoTelp   string `json:"no_telp" form:"no_telp"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_new_password"`
}

type UserVolunteerHistory struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	ImageURL string    `json:"image_url"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

type UserVolunteerHistoryDetail struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	ImageURL        string `json:"image_url"`
	Location        string `json:"location"`
	ContentActivity string `json:"content_activity"`
}

type UserDonationHistory struct {
	ID       uint   `json:"id"`
	Tittle   string `json:"tittle"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
	Amount   int    `json:"amount"`
}

type UserFundraisingsResponse struct {
	ID              uint   `json:"id"`
	FundraisingID   uint   `json:"fundraising_id"`
	ImageUrl        string `json:"image_url"`
	Title           string `json:"title"`
	CategoryName    string `json:"category_name"`
	CurrentProgress int    `json:"current_progress"`
	TargetAmount    int    `json:"target_amount"`
	EndDate         string `json:"end_date"`
}

func ToUserFundraisingsResponse(fundraisingBookmark entities.UserBookmarkFundraising) UserFundraisingsResponse {
	return UserFundraisingsResponse{
		ID:              fundraisingBookmark.ID,
		FundraisingID:   fundraisingBookmark.FundraisingID,
		ImageUrl:        fundraisingBookmark.Fundraising.ImageUrl,
		Title:           fundraisingBookmark.Fundraising.Title,
		CategoryName:    fundraisingBookmark.Fundraising.FundraisingCategory.Name,
		CurrentProgress: fundraisingBookmark.Fundraising.CurrentProgress,
		TargetAmount:    fundraisingBookmark.Fundraising.GoalAmount,
		EndDate:         fundraisingBookmark.Fundraising.EndDate.Format("2006-01-02"),
	}
}

func ToAllUserFundraisingsResponse(fundraisings []entities.UserBookmarkFundraising) []UserFundraisingsResponse {
	var result []UserFundraisingsResponse
	for _, fundraising := range fundraisings {
		result = append(result, ToUserFundraisingsResponse(fundraising))
	}
	return result
}

type UserArticleBookmark struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	Date      string `json:"date"`
}

func ToUserArticleBookmarkResponse(articleBookmark entities.UserBookmarkArticle) UserArticleBookmark {
	return UserArticleBookmark{
		ID:        articleBookmark.ID,
		ArticleID: articleBookmark.ArticleID,
		Title:     articleBookmark.Article.Title,
		ImageURL:  articleBookmark.Article.ImageURL,
		Date:      articleBookmark.Article.CreatedAt.Format("2006-01-02"),
	}
}

func ToAllUserArticleBookmarkResponse(articles []entities.UserBookmarkArticle) []UserArticleBookmark {
	var result []UserArticleBookmark
	for _, article := range articles {
		result = append(result, ToUserArticleBookmarkResponse(article))
	}
	return result
}
