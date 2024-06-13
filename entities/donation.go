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

type DonationComment struct {
	gorm.Model
	DonationID uint     `json:"-"`
	Donation   Donation `json:"donation" gorm:"foreignKey:DonationID"`
	Comment    string   `json:"comment" gorm:"type:varchar(255)"`
	TotalLikes int      `json:"total_likes" gorm:"type:int"`
}

type LikeDonationComment struct {
	gorm.Model
	DonationCommentID uint            `json:"-"`
	DonationComment   DonationComment `json:"donation_comment" gorm:"foreignKey:DonationCommentID"`
	UserID            uint            `json:"-"`
	User              User            `json:"user" gorm:"foreignKey:UserID"`
}

