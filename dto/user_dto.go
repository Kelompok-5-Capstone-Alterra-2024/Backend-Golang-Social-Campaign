package dto

type RegisterRequest struct {
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
}
