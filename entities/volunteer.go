package entities

import (
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	OrganizationID       uint      `json:"organization_id" gorm:"type:bigint"`
	Title                string    `json:"title" gorm:"type:varchar(255)"`
	ContentActivity      string    `json:"content_activity" gorm:"type:text"`
	Location             string    `json:"location" gorm:"type:varchar(255)"`
	Date                 time.Time `json:"date"`
	TargetVolunteer      int       `json:"target_volunteer"`
	RegisteredVolunteer  int       `json:"registered_volunteer"`
	RegistrationDeadline time.Time `json:"registration_deadline"`
	ImageURL             string    `json:"image_url" gorm:"type:varchar(255)"`
}
