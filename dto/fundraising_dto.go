package dto

import (
	"capstone/entities"
	"time"
)

type CreateFundraisingRequest struct {
	ImageUrl       string `json:"image_url"`
	Title          string `json:"title"`
	TargetAmount   int    `json:"target_amount"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Description    string `json:"description"`
	CategoryID     uint   `json:"category_id"`
	OrganizationID uint   `json:"organization_id"`
}

type FundraisingsResponse struct {
	ID              uint   `json:"id"`
	ImageUrl        string `json:"image_url"`
	Title           string `json:"title"`
	CategoryName    string `json:"category_name"`
	CurrentProgress int    `json:"current_progress"`
	EndDate         string `json:"end_date"`
}

func ToFundraisingsResponse(fundraising entities.Fundraising) FundraisingsResponse {
	return FundraisingsResponse{
		ID:              fundraising.ID,
		ImageUrl:        fundraising.ImageUrl,
		Title:           fundraising.Title,
		CategoryName:    fundraising.FundraisingCategory.Name,
		CurrentProgress: fundraising.CurrentProgress,
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
	ImageUrl         string                       `json:"image_url"`
	Title            string                       `json:"title"`
	GoalAmount       int                          `json:"goal_amount"`
	CurrentProgress  int                          `json:"current_progress"`
	EndDate          string                       `json:"end_date"`
	OrganizationName string                       `json:"organization_name"`
	OrgIsVerified    bool                         `json:"org_is_verified"`
	Description      string                       `json:"description"`
	UserDonated      UserDonatedResponse          `json:"user_donated"`
	Comment          []FundraisingCommentResponse `json:"comment"`
}

type UserDonatedResponse struct {
	UserAvatarDonatedResponse UserAvatarDonatedResponse `json:"user_avatar_donated"`
	TotalUserDonated          int                       `json:"total_user_donated"`
}

type UserAvatarDonatedResponse struct {
	UserID uint   `json:"user_id"`
	Avatar string `json:"avatar"`
}

type FundraisingCommentResponse struct {
	UserCommentResponse UserCommentResponse `json:"user_comment"`
	Body                string              `json:"body"`
	TotalLikes          int                 `json:"total_likes"`
	CreatedAt           string              `json:"created_at"`
}

type UserCommentResponse struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

func ToFundraisingResponse(fundraising entities.Fundraising, comments []entities.DonationComment, donations []entities.Donation) FundraisingResponse {

	uniqueUserAvatars := make(map[uint]string)
	for _, donation := range donations {
		uniqueUserAvatars[donation.UserID] = donation.User.Avatar
	}

	// Get the avatar of the first unique user who donated
	var userAvatarDonatedResponse UserAvatarDonatedResponse
	for userID, avatar := range uniqueUserAvatars {
		userAvatarDonatedResponse = UserAvatarDonatedResponse{
			UserID: userID,
			Avatar: avatar,
		}
		break
	}

	userDonatedResponse := UserDonatedResponse{
		UserAvatarDonatedResponse: userAvatarDonatedResponse,
		TotalUserDonated:          len(uniqueUserAvatars),
	}

	commentResponses := make([]FundraisingCommentResponse, len(comments))
	for i, comment := range comments {
		commentResponses[i] = FundraisingCommentResponse{
			UserCommentResponse: UserCommentResponse{
				Avatar:   comment.Donation.User.Avatar,
				Username: comment.Donation.User.Username,
			},
			Body:       comment.Comment,
			TotalLikes: comment.TotalLikes,
			CreatedAt:  comment.CreatedAt.Format(time.RFC3339),
		}
	}

	return FundraisingResponse{
		ImageUrl:         fundraising.ImageUrl,
		Title:            fundraising.Title,
		GoalAmount:       fundraising.GoalAmount,
		CurrentProgress:  fundraising.CurrentProgress,
		EndDate:          fundraising.EndDate.Format("2006-01-02"),
		OrganizationName: fundraising.Organization.Name,
		OrgIsVerified:    fundraising.Organization.IsVerified,
		Description:      fundraising.Description,
		UserDonated:      userDonatedResponse,
		Comment:          commentResponses,
	}
}
