package dto

import (
	"capstone/entities"
	"time"
)

type OrganizationRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	IsVerified  bool   `json:"is_verified" form:"is_verified" validate:"required"`
	StartDate   string `json:"start_date" form:"start_date" validate:"required"`
	Website     string `json:"website" form:"website" validate:"required"`
	Instagram   string `json:"instagram" form:"instagram" validate:"required"`
	NoRekening  string `json:"no_rek" form:"no_rek" validate:"required"`
}

type OrganizationResponse struct {
	ID          uint                   `json:"id" `
	Name        string                 `json:"name" `
	CreatedAt   time.Time              `json:"created_at" `
	Avatar      string                 `json:"avatar" `
	Description string                 `json:"description" `
	IsVerified  bool                   `json:"is_verified" `
	Fundraising []FundraisingsResponse `json:"fundraising" `
}

func ToOrganizationResponse(organization entities.Organization, fundraisings []entities.Fundraising) OrganizationResponse {

	var Fundraisings []FundraisingsResponse
	for _, fundraising := range fundraisings {
		Fundraisings = append(Fundraisings, ToFundraisingsResponse(fundraising))
	}

	return OrganizationResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		CreatedAt:   organization.CreatedAt,
		Avatar:      organization.Avatar,
		IsVerified:  organization.IsVerified,
		Description: organization.Description,
		Fundraising: Fundraisings,
	}
}
