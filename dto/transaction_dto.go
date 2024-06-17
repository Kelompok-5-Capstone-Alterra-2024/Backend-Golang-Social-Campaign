package dto

import "capstone/entities"

type TransactionHistoryRespone struct {
	ID               uint   `json:"id" `
	OrganizationName string `json:"organization_name"`
	Amount           int    `json:"amount"`
	PaymentMethod    string `json:"payment_method"`
	NoRekening       string `json:"no_rekening"`
	CreatedAt        string `json:"created_at"`
}

func ToTransactionHistoryRespone(transaction entities.Transaction) TransactionHistoryRespone {
	return TransactionHistoryRespone{
		ID:               transaction.ID,
		OrganizationName: transaction.Fundraising.Organization.Name,
		Amount:           transaction.Amount,
		PaymentMethod:    transaction.BankName,
		NoRekening:       transaction.NoRekening,
		CreatedAt:        transaction.CreatedAt.Format("2006-01-02"),
	}
}

func ToTransactionHistoriesRespone(transactions []entities.Transaction) []TransactionHistoryRespone {
	var result []TransactionHistoryRespone
	for _, transaction := range transactions {
		result = append(result, ToTransactionHistoryRespone(transaction))
	}
	return result
}
