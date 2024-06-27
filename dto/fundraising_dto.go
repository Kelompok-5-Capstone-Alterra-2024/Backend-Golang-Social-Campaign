package dto

import (
	"capstone/entities"
	"time"
)

type CreateFundraisingRequest struct {
	ImageUrl       string `json:"image_url" form:"image_url" required:"true"`
	Title          string `json:"title" form:"title" required:"true"`
	TargetAmount   int    `json:"target_amount" form:"target_amount" required:"true"`
	StartDate      string `json:"start_date" form:"start_date"`
	EndDate        string `json:"end_date" form:"end_date"`
	Description    string `json:"description" form:"description" required:"true"`
	CategoryID     uint   `json:"category_id" form:"category_id" required:"true"`
	OrganizationID uint   `json:"organization_id" form:"organization_id" required:"true"`
}

type FundraisingsResponse struct {
	ID              uint   `json:"id"`
	ImageUrl        string `json:"image_url"`
	Title           string `json:"title"`
	CategoryName    string `json:"category_name"`
	CurrentProgress int    `json:"current_progress"`
	TargetAmount    int    `json:"target_amount"`
	EndDate         string `json:"end_date"`
}

func ToFundraisingsResponse(fundraising entities.Fundraising) FundraisingsResponse {
	return FundraisingsResponse{
		ID:              fundraising.ID,
		ImageUrl:        fundraising.ImageUrl,
		Title:           fundraising.Title,
		CategoryName:    fundraising.FundraisingCategory.Name,
		CurrentProgress: fundraising.CurrentProgress,
		TargetAmount:    fundraising.GoalAmount,
		EndDate:         fundraising.EndDate.Format("2006-01-02"),
	}
}

func ToAllFundraisingsResponse(fundraisings []entities.Fundraising) []FundraisingsResponse {
	var result []FundraisingsResponse
	for _, fundraising := range fundraisings {
		result = append(result, ToFundraisingsResponse(fundraising))
	}
	return result
}

type FundraisingResponse struct {
	ID              uint                         `json:"id"`
	Organization    FundraisingOrg               `json:"organization"`
	ImageUrl        string                       `json:"image_url"`
	Title           string                       `json:"title"`
	GoalAmount      int                          `json:"goal_amount"`
	CurrentProgress int                          `json:"current_progress"`
	EndDate         string                       `json:"end_date"`
	Description     string                       `json:"description"`
	UserDonated     UserDonatedResponse          `json:"user_donated"`
	Comment         []FundraisingCommentResponse `json:"comment"`
}

type FundraisingOrg struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	IsVerified bool   `json:"is_verified"`
}

type UserDonatedResponse struct {
	UserAvatarDonatedResponse []UserAvatarDonatedResponse `json:"user_avatar_donated"`
	TotalUserDonated          int                         `json:"total_user_donated"`
}

type UserAvatarDonatedResponse struct {
	UserID uint   `json:"user_id"`
	Avatar string `json:"avatar"`
}

type FundraisingCommentResponse struct {
	ID                  uint                `json:"id"`
	UserCommentResponse UserCommentResponse `json:"user_comment"`
	Body                string              `json:"body"`
	TotalLikes          int                 `json:"total_likes"`
	CreatedAt           string              `json:"created_at"`
}

type UserCommentResponse struct {
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

func ToFundraisingResponse(fundraising entities.Fundraising, comments []entities.DonationManualComment, donations []entities.DonationManual) FundraisingResponse {

	uniqueUserAvatars := make(map[uint]string)
	for _, donation := range donations {
		uniqueUserAvatars[donation.UserID] = donation.User.Avatar
	}

	// Get the avatar of the first four unique user who donated
	userAvatarDonatedResponse := []UserAvatarDonatedResponse{}
	for userID, avatar := range uniqueUserAvatars {
		if len(userAvatarDonatedResponse) == 4 {
			break
		}
		userAvatarDonatedResponse = append(userAvatarDonatedResponse, UserAvatarDonatedResponse{
			UserID: userID,
			Avatar: avatar,
		})
	}

	userDonatedResponse := UserDonatedResponse{
		UserAvatarDonatedResponse: userAvatarDonatedResponse,
		TotalUserDonated:          len(uniqueUserAvatars),
	}

	fundraisingOrg := FundraisingOrg{
		ID:         fundraising.Organization.ID,
		Name:       fundraising.Organization.Name,
		Avatar:     fundraising.Organization.Avatar,
		IsVerified: fundraising.Organization.IsVerified,
	}

	commentResponses := make([]FundraisingCommentResponse, len(comments))
	for i, comment := range comments {
		commentResponses[i] = FundraisingCommentResponse{
			ID: comment.ID,
			UserCommentResponse: UserCommentResponse{
				UserID:   comment.DonationManual.User.ID,
				Avatar:   comment.DonationManual.User.Avatar,
				Username: comment.DonationManual.User.Username,
			},
			Body:       comment.Comment,
			TotalLikes: comment.TotalLikes,
			CreatedAt:  comment.CreatedAt.Format(time.RFC3339),
		}
	}

	return FundraisingResponse{
		ID:              fundraising.ID,
		ImageUrl:        fundraising.ImageUrl,
		Title:           fundraising.Title,
		GoalAmount:      fundraising.GoalAmount,
		CurrentProgress: fundraising.CurrentProgress,
		EndDate:         fundraising.EndDate.Format("2006-01-02"),
		Organization:    fundraisingOrg,
		Description:     fundraising.Description,
		UserDonated:     userDonatedResponse,
		Comment:         commentResponses,
	}
}
