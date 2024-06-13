package entities

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name        string    `json:"name" form:"name" gorm:"type:varchar(255)"`
	Description string    `json:"description" form:"description" gorm:"type:text"`
	Avatar      string    `json:"avatar" form:"avatar" gorm:"type:varchar(255)"`
	IsVerified  bool      `json:"is_verified" form:"is_verified" gorm:"type:bool"`
	StartDate   time.Time `json:"start_date" form:"start_date" gorm:"type:date"`
	Contact     string    `json:"contact" form:"contact" gorm:"type:varchar(255)"`
	Website     string    `json:"website" form:"website" gorm:"type:varchar(255)"`
	Instagram   string    `json:"instagram" form:"instagram" gorm:"type:varchar(255)"`
	NoRekening  string    `json:"no_rek" form:"no_rek" gorm:"type:varchar(255)"`
}
