package dto

import (
	"capstone/entities"
)

type DonationRequest struct {
	ID     uint          `json:"id" `
	Amount int           `json:"amount" form:"amount"`
	User   entities.User `json:"user"`
}
