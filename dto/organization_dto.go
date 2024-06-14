package dto

import (
	"capstone/entities"
)

type OrganizationRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	IsVerified  bool   `json:"is_verified" form:"is_verified" validate:"required"`
	JoinDate    string `json:"join_date" form:"join_date" validate:"required"`
	Website     string `json:"website" form:"website" validate:"required"`
	Instagram   string `json:"instagram" form:"instagram" validate:"required"`
	NoRekening  string `json:"no_rek" form:"no_rek" validate:"required"`
}

type OrganizationFundraisingResponse struct {
	ID          uint                   `json:"id" `
	Name        string                 `json:"name" `
	JoinDate    string                 `json:"join_date" `
	Avatar      string                 `json:"avatar" `
	Description string                 `json:"description" `
	IsVerified  bool                   `json:"is_verified" `
	Website     string                 `json:"website" `
	Instagram   string                 `json:"instagram" `
	Fundraising []FundraisingsResponse `json:"fundraising" `
}

func ToOrganizationFundraisingResponse(organization entities.Organization, fundraisings []entities.Fundraising) OrganizationFundraisingResponse {

	var Fundraisings []FundraisingsResponse
	for _, fundraising := range fundraisings {
		Fundraisings = append(Fundraisings, ToFundraisingsResponse(fundraising))
	}

	return OrganizationFundraisingResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		JoinDate:    organization.StartDate.Format("2006-01-02"),
		Avatar:      organization.Avatar,
		IsVerified:  organization.IsVerified,
		Description: organization.Description,
		Website:     organization.Website,
		Instagram:   organization.Instagram,
		Fundraising: Fundraisings,
	}
}

type OrganizationVolunteerResponse struct {
	ID          uint                  `json:"id" `
	Name        string                `json:"name" `
	JoinDate    string                `json:"join_date" `
	Avatar      string                `json:"avatar" `
	Description string                `json:"description" `
	IsVerified  bool                  `json:"is_verified" `
	Website     string                `json:"website" `
	Instagram   string                `json:"instagram" `
	Volunteers  []VolunteersResponses `json:"volunteers" `
}

func ToOrganizationVolunteerResponse(organization entities.Organization, volunteers []entities.Volunteer) OrganizationVolunteerResponse {

	var Volunteers []VolunteersResponses

	for _, volunteer := range volunteers {
		Volunteers = append(Volunteers, ToVolunteersResponses(volunteer))
	}

	return OrganizationVolunteerResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		JoinDate:    organization.StartDate.Format("2006-01-02"),
		Avatar:      organization.Avatar,
		IsVerified:  organization.IsVerified,
		Description: organization.Description,
		Website:     organization.Website,
		Instagram:   organization.Instagram,
		Volunteers:  Volunteers,
	}
}
