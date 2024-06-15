package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname   string `json:"fullname" gorm:"type:varchar(255)"`
	Username   string `json:"username" gorm:"type:varchar(255)"`
	Email      string `json:"email" gorm:"type:varchar(255)"`
	Password   string `json:"password" gorm:"type:varchar(255)"`
	Avatar     string `json:"avatar" gorm:"type:varchar(255)"`
	NoTelp     string `json:"no_telp"`
	Token      string `gorm:"-"`
	ResetToken string `gorm:"type:varchar(255)"`
	OTP        string `gorm:"type:varchar(255)"`
}

type UserBookmarkVolunteerVacancy struct {
	gorm.Model
	UserID               uint      `json:"-"`
	User                 User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	VolunteerVacanciesID uint      `json:"-"`
	Volunteer            Volunteer `json:"volunteer" gorm:"foreignKey:VolunteerVacanciesID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type UserBookmarkFundraising struct {
	gorm.Model
	UserID        uint        `json:"-"`
	User          User        `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	FundraisingID uint        `json:"-"`
	Fundraising   Fundraising `json:"fundraising" gorm:"foreignKey:FundraisingID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type UserBookmarkArticle struct {
	gorm.Model
	UserID    uint    `json:"-"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	ArticleID uint    `json:"-"`
	Article   Article `json:"article" gorm:"foreignKey:ArticleID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// type OTP struct {
// 	gorm.Model
// 	UserID     int    `json:"user_id" gorm:"index;unique"`
// 	User       User   `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
// 	OTP        string `json:"otp" gorm:"type:varchar(255)"`
// 	ExpiredOTP int64  `json:"expired_otp" gorm:"type:bigint"`
// }
