package entities

type PaymentTransaction struct {
	ID              int    `json:"id"`
	Amount          int    `json:"amount"`
	FundraisingName string `json:"fundraising_name"`
}
