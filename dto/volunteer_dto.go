package dto

import (
	"capstone/entities"
	"fmt"
	"time"
)

type VolunteerRequest struct {
	OrganizationID       uint   `json:"organization_id"`
	Title                string `json:"title"`
	ContentActivity      string `json:"content_activity"`
	Location             string `json:"location"`
	Date                 string `json:"date"`
	TargetVolunteer      int    `json:"target_volunteer"`
	RegisteredVolunteer  int    `json:"registered_volunteer"`
	RegistrationDeadline string `json:"registration_deadline"`
	ImageURL             string `json:"image_url"`
}

func (r *VolunteerRequest) ToEntity() (entities.Volunteer, error) {
	date, err := time.ParseInLocation("02/01/2006", r.Date, time.FixedZone("GMT+7", 7*60*60))
	if err != nil {
		return entities.Volunteer{}, fmt.Errorf("invalid date format: %v", err)
	}

	registrationDeadline, err := time.ParseInLocation("02/01/2006", r.RegistrationDeadline, time.FixedZone("GMT+7", 7*60*60))
	if err != nil {
		return entities.Volunteer{}, fmt.Errorf("invalid registration deadline format: %v", err)
	}

	return entities.Volunteer{
		OrganizationID:       r.OrganizationID,
		Title:                r.Title,
		ContentActivity:      r.ContentActivity,
		Location:             r.Location,
		Date:                 date,
		TargetVolunteer:      r.TargetVolunteer,
		RegisteredVolunteer:  r.RegisteredVolunteer,
		RegistrationDeadline: registrationDeadline,
		ImageURL:             r.ImageURL,
	}, nil
}
