package entities

import (
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	Title                string    `json:"title" gorm:"type:varchar(255)"`
	ContentActivity      string    `json:"content_activity" gorm:"type:text"`
	Location             string    `json:"location" gorm:"type:varchar(255)"`
	Date                 time.Time `json:"date" gorm:"type:datetime"`
	TargetVolunteer      int       `json:"target_volunteer" gorm:"type:int"`
	RegisteredVolunteer  int       `json:"registered_volunteer" gorm:"type:int"`
	RegistrationDeadline time.Time `json:"registration_deadline" gorm:"type:datetime"`
	Image                string    `json:"image" gorm:"type:varchar(255)"`
}
