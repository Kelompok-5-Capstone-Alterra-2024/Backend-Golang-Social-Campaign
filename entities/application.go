package entities

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	IgImageURL         string    `json:"ig_image_url" gorm:"type:varchar(255)"`
	YtImageURL         string    `json:"yt_image_url" gorm:"type:varchar(255)"`
	UserID             uint      `json:"user_id" gorm:"type:bigint"`
	User               User      `json:"user" gorm:"foreignKey:UserID"`
	VolunteerVacancyID uint      `json:"vacancy_id" gorm:"type:bigint"`
	Volunteer          Volunteer `json:"volunteer" gorm:"foreignKey:VolunteerVacancyID"`
	Reason             string    `json:"reason" gorm:"type:varchar(255)"`
	Age                int       `json:"age" gorm:"type:int"`
	Job                string    `json:"job" gorm:"type:varchar(255)"`
}
