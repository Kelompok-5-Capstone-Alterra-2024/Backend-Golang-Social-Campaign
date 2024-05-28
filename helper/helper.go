package helper

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

type generalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GeneralResponse(status, message string) generalResponse {
	messageRes := generalResponse{
		Status:  status,
		Message: message,
	}
	return messageRes
}

type responseWithData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(status, message string, data interface{}) responseWithData {
	messageRes := responseWithData{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return messageRes
}

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(status, message string, err any) errorResponse {
	messageRes := errorResponse{
		Status:  status,
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

func SendTokenRestPassword(email string, token string) error {

	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"hanggoroseto8@gmail.com",
		"pcxf rviq wvfz nfyy",
	)

	resetURL := fmt.Sprintf("%s/reset-password?token=%s", "http://localhost:8080", token)

	m := gomail.NewMessage()
	m.SetHeader("From", "hanggoroseto8@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")
	m.SetBody("text/plain", "Click the link to reset your password: "+resetURL)

	return dialer.DialAndSend(m)
}

func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
