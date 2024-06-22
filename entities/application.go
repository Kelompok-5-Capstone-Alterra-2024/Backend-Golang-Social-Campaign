package entities

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	IgImageURL string    `json:"ig_image_url" gorm:"type:varchar(255)"`
	YtImageURL string    `json:"yt_image_url" gorm:"type:varchar(255)"`
	UserID     uint      `json:"-"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	VacancyID  uint      `json:"-"`
	Volunteer  Volunteer `json:"volunteer" gorm:"foreignKey:VacancyID"`
	Job        string    `json:"job" gorm:"type:varchar(255)"`
	Reason     string    `json:"reason" gorm:"type:varchar(255)"`
	Age        string    `json:"age" gorm:"type:varchar(255)"`
}
