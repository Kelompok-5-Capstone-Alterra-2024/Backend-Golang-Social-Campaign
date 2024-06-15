package entities

import (
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	Title                string       `json:"title" gorm:"type:varchar(255)"`
	OrganizationID       uint         `json:"-"`
	Organization         Organization `json:"organization" gorm:"foreignKey:OrganizationID"`
	ContentActivity      string       `json:"content_activity" gorm:"type:text"`
	Location             string       `json:"location" gorm:"type:varchar(255)"`
	StartDate            time.Time    `json:"start_date" gorm:"datetime"`
	EndDate              time.Time    `json:"end_date" gorm:"datetime"`
	TargetVolunteer      int          `json:"target_volunteer"`
	Status               string       `json:"status" gorm:"type:enum('active','inactive')"`
	RegisteredVolunteer  int          `json:"registered_volunteer"`
	RegistrationDeadline time.Time    `json:"registration_deadline" gorm:"datetime"`
	ImageURL             string       `json:"image_url" gorm:"type:varchar(255)"`
}
