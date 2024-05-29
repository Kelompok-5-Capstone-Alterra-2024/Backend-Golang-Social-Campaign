package dto

import (
	"capstone/entities"
	"time"
)

type VolunteerRequest struct {
	Title                string `json:"title"`
	ContentActivity      string `json:"content_activity"`
	Location             string `json:"location"`
	Date                 string `json:"date"`
	TargetVolunteer      int    `json:"target_volunteer"`
	RegisteredVolunteer  int    `json:"registered_volunteer"`
	RegistrationDeadline string `json:"registration_deadline"`
	Image                string `json:"image"`
}

func (req *VolunteerRequest) ToEntity() entities.Volunteer {
	dateTimeLayout := "02/01/2006 15:04:05 -0700"

	date, err := time.Parse(dateTimeLayout, req.Date)
	if err != nil {
		return entities.Volunteer{}
	}

	deadline, err := time.Parse(dateTimeLayout, req.RegistrationDeadline)
	if err != nil {
		return entities.Volunteer{}
	}

	return entities.Volunteer{
		Title:                req.Title,
		ContentActivity:      req.ContentActivity,
		Location:             req.Location,
		Date:                 date,
		TargetVolunteer:      req.TargetVolunteer,
		RegisteredVolunteer:  req.RegisteredVolunteer,
		RegistrationDeadline: deadline,
		Image:                req.Image,
	}
}
