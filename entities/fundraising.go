package entities

import (
	"time"

	"gorm.io/gorm"
)

type Fundraising struct {
	gorm.Model
	FundraisingCategoryID uint                `json:"-"`
	FundraisingCategory   FundraisingCategory `json:"category" gorm:"foreignKey:FundraisingCategoryID"`
	OrganizationID        uint                `json:"-"`
	Organization          Organization        `json:"organization" gorm:"foreignKey:OrganizationID"`
	Title                 string              `json:"title" gorm:"type:varchar(255)"`
	ImageUrl              string              `json:"image_url" gorm:"type:varchar(255)"`
	Description           string              `json:"description" gorm:"type:varchar(255)"`
	Status                interface{}         `json:"status" gorm:"type:varchar(255)"`
	GoalAmount            int                 `json:"goal_amount" gorm:"type:int"`
	CurrentProgress       int                 `json:"current_progress" gorm:"type:int"`
	StartDate             time.Time           `json:"start_date" gorm:"datetime"`
	EndDate               time.Time           `json:"end_date" gorm:"datetime"`
}

type FundraisingCategory struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ImageUrl string `json:"image_url" gorm:"type:varchar(255)"`
}
