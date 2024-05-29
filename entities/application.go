package entities

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	VolunteerID uint   `json:"volunteer_id" gorm:"type:int"`
	UserID      uint   `json:"user_id" gorm:"type:int"`
	Status      string `json:"status" gorm:"type:varchar(255)"`
	IgImage     string `json:"ig_image" gorm:"type:varchar(255)"`
}
