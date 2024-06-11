package entities

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	UserID        uint        `json:"-"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	Amount        int         `json:"amount" gorm:"type:int"`
	Status        string      `json:"status" gorm:"type:varchar(255)"`
	Code          string      `json:"code" gorm:"type:varchar(255)"`
	FundraisingID uint        `json:"-"`
	Fundraising   Fundraising `json:"fundraising" gorm:"foreignKey:FundraisingID"`
	PaymentUrl    string      `json:"payment_url" gorm:"type:varchar(255)"`
	PaymentMethod string      `json:"payment_method" gorm:"type:varchar(255)"`
}
