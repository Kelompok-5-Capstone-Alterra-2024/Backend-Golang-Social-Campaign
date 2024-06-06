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
	Contact     string `json:"contact" form:"contact" validate:"required"`
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
