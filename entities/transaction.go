package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount        int         `json:"amount" gorm:"type:int"`
	BankName      string      `json:"bank_name" gorm:"type:varchar(255)"`
	NoRekening    string      `json:"no_rekening" gorm:"type:varchar(255)"`
	Name          string      `json:"name" gorm:"type:varchar(255)"`
	ImagePayment  string      `json:"image_payment" gorm:"type:varchar(255)"`
	FundraisingID uint        `json:"-"`
	Fundraising   Fundraising `json:"fundraising" gorm:"foreignKey:FundraisingID"`
}
