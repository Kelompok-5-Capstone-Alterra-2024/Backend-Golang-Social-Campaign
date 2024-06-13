package entities

import "gorm.io/gorm"

type DonationManual struct {
	gorm.Model
	UserID        uint        `json:"-"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	Amount        int         `json:"amount" gorm:"type:int"`
	Status        string      `json:"status" gorm:"type:varchar(255)"`
	FundraisingID uint        `json:"-"`
	Fundraising   Fundraising `json:"fundraising" gorm:"foreignKey:FundraisingID"`
	ImagePayment  string      `json:"image_payment" gorm:"type:varchar(255)"`
}

type DonationManualComment struct {
	gorm.Model
	DonationManualID uint           `json:"-"`
	DonationManual   DonationManual `json:"donation" gorm:"foreignKey:DonationManualID"`
	Comment          string         `json:"comment" gorm:"type:varchar(255)"`
	TotalLikes       int            `json:"total_likes" gorm:"type:int"`
}

type LikeDonationManualComment struct {
	gorm.Model
	DonationManualCommentID uint                  `json:"-"`
	DonationManualComment   DonationManualComment `json:"donation_comment" gorm:"foreignKey:DonationManualCommentID"`
	UserID                  uint                  `json:"-"`
	User                    User                  `json:"user" gorm:"foreignKey:UserID"`
}
