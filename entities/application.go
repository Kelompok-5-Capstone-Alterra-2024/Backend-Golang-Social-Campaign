package entities

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	IgImageURL string `json:"ig_image_url" gorm:"type:varchar(255)"`
	YtImageURL string `json:"yt_image_url" gorm:"type:varchar(255)"`
	CustomerID uint   `json:"customer_id" gorm:"type:bigint"`
	VacancyID  uint   `json:"vacancy_id" gorm:"type:bigint"`
	Reason     string `json:"reason" gorm:"type:varchar(255)"`
	Age        int    `json:"age" gorm:"type:int"`
	Job        string `json:"job" gorm:"type:varchar(255)"`
}
