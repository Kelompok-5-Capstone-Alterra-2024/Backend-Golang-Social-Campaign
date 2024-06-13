package entities

import (
	"gorm.io/gorm"
)

type TestimoniVolunteer struct {
	gorm.Model
	UserID    uint      `json:"customer_id" gorm:"type:bigint"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	VacancyID uint      `json:"-"`
	Volunteer Volunteer `json:"volunteer" gorm:"foreignKey:VacancyID"`
	Testimoni string    `json:"testimoni" gorm:"type:varchar(255)"`
	Rating    string    `json:"rating" gorm:"type:enum('1','2','3','4','5')"`
}
