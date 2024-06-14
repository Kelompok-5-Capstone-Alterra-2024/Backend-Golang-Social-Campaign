package helper

import (
	// "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type generalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func GeneralResponse(success bool, message string) generalResponse {
	messageRes := generalResponse{
		Success: success,
		Message: message,
	}
	return messageRes
}

type responseWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(success bool, message string, data interface{}) responseWithData {
	messageRes := responseWithData{
		Success: success,
		Message: message,
		Data:    data,
	}

	return messageRes
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(success bool, message string, err any) errorResponse {
	messageRes := errorResponse{
		Success: success,
		Message: message,
		Error:   err,
	}

	return messageRes
}

func GetToken(auth string) string {
	splittedAuth := strings.Split(auth, "Bearer ")
	return splittedAuth[1]
}

func DecodePayload(token string) (map[string]interface{}, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT token format")
	}

	// Decode base64 encoded payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		return nil, err
	}

	return payloadMap, nil
}

func GetUserIDFromJWT(c echo.Context) (int, error) {
	authorization := c.Request().Header.Get("Authorization")
	if authorization == "" {
		return 0, errors.New("unauthorized")
	}

	jwtToken := GetToken(authorization)

	jwt_payload, err := DecodePayload(jwtToken)
	if err != nil {
		return 0, err
	}

	// Get user id from jwt payload
	user_id, ok := jwt_payload["id"].(float64)
	if !ok {
		return 0, errors.New("unauthorized")
	}

	return int(user_id), nil
}

func SendTokenRestPassword(email string, token string) error {

	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"hanggoroseto8@gmail.com",
		"pcxf rviq wvfz nfyy",
	)

	resetURL := fmt.Sprintf("%s/reset-password?token=%s", "https://capstone-alterra-424313.as.r.appspot.com/api/v1", token)

	m := gomail.NewMessage()
	m.SetHeader("From", "hanggoroseto8@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	m.SetBody("text/plain", "Click the link to reset your password: "+resetURL)

	return dialer.DialAndSend(m)
}

func SendOtpResetPassword(email string, otp string) error {
	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"hanggoroseto8@gmail.com",
		"pcxf rviq wvfz nfyy",
	)

	m := gomail.NewMessage()
	m.SetHeader("From", "hanggoroseto8@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	m.SetBody("text/plain", "Your OTP is: "+otp)

	return dialer.DialAndSend(m)
}

func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func GenerateRandomOTP(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	const n = "0123456789"

	otp := make([]byte, length)
	for i := range otp {
		otp[i] = n[r.Intn(len(n))]
	}

	return string(otp)
}
