package dto

import "capstone/entities"

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
	ImageUrl         string `json:"image_url"`
	Title            string `json:"title"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentProgress  int    `json:"current_progress"`
	EndDate          string `json:"end_date"`
	OrganizationName string `json:"organization_name"`
	OrgIsVerified    bool   `json:"org_is_verified"`
	Description      string `json:"description"`
}

func ToFundraisingResponse(fundraising entities.Fundraising) FundraisingResponse {
	return FundraisingResponse{
		ImageUrl:         fundraising.ImageUrl,
		Title:            fundraising.Title,
		GoalAmount:       fundraising.GoalAmount,
		CurrentProgress:  fundraising.CurrentProgress,
		EndDate:          fundraising.EndDate.Format("2006-01-02"),
		OrganizationName: fundraising.Organization.Name,
		OrgIsVerified:    fundraising.Organization.IsVerified,
		Description:      fundraising.Description,
	}
}
