package helper

import (
	"capstone/entities"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
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

type DataResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewDataResponse(status string, data interface{}) *DataResponse {
	return &DataResponse{
		Status:  status,
		Message: "success",
		Data:    data,
	}
}

type PaginationResponse struct {
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	TotalRecords int64       `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	PageSize     int         `json:"page_size"`
}

func ResponseWithPagination(status, message string, data interface{}, page, limit int, totalRecords int64) PaginationResponse {
	totalPages := int((totalRecords + int64(limit) - 1) / int64(limit))
	return PaginationResponse{
		Status:       status,
		Message:      message,
		Data:         data,
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     limit,
	}
}

func StringToUint(s string) (uint, error) {
	id, err := strconv.ParseUint(strings.TrimSpace(s), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
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

func GenerateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func GenerateRandomOTP(otpLent int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	const n = "0123456789"

	otp := make([]byte, otpLent)
	for i := range otp {
		otp[i] = n[r.Intn(len(n))]
	}

	return string(otp)
}

func GetPaymentUrl(donation entities.PaymentTransaction, user entities.User) (string, error) {
	midClient := midtrans.NewClient()
	server := "SB-Mid-server-x_R3_BBoJmSU_bRRxcBWV9pg"
	client := "SB-Mid-client-YStDTAnO_VeyBKdH"
	midClient.ServerKey = server
	midClient.ClientKey = client
	midClient.APIEnvType = midtrans.Sandbox
	orderID := donation.ID
	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Fullname,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(orderID),
			GrossAmt: int64(donation.Amount),
		},
		Items: &[]midtrans.ItemDetail{
			{
				Name:  donation.FundraisingName,
				Qty:   1,
				Price: int64(donation.Amount),
			},
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil

}
