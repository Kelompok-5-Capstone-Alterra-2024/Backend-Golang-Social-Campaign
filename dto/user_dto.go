package dto

type RegisterRequest struct {
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	NoTelp      string `json:"no_telp"`
	Password    string `json:"password"`
	ConfirmPass string `json:"confirm_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Password    string `json:"new_password"`
	ConfirmPass string `json:"confirm_password"`
}
