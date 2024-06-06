package entities

import (
	"gorm.io/gorm"
)

type TestimoniVolunteer struct {
	gorm.Model
	CustomerID  uint   `json:"customer_id" gorm:"type:bigint"`
	VolunteerID uint   `json:"volunteer_id" gorm:"type:bigint"`
	Testimoni   string `json:"testimoni" gorm:"type:varchar(255)"`
	Rating      string `json:"rating" gorm:"type:enum('1','2','3','4','5')"`
}
