package service

import (
	"capstone/dto"
	"capstone/entities"
	"capstone/repositories"
	"fmt"
)

type TransactionService interface {
	CreateTransaction(transaction dto.DistributeFundFundraisingRequest) (entities.Transaction, error)
	GetTransactionByID(id uint) (entities.Transaction, error)
	GetTransactions(limit int, offset int) ([]entities.Transaction, error)
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
	adminRepo             repositories.AdminRepository
}

func NewTransactionService(transactionRepository repositories.TransactionRepository, adminRepo repositories.AdminRepository) *transactionService {
	return &transactionService{transactionRepository, adminRepo}
}

func (s *transactionService) CreateTransaction(transaction dto.DistributeFundFundraisingRequest) (entities.Transaction, error) {
	transactionEntity := entities.Transaction{
		Amount:        transaction.Amount,
		BankName:      transaction.BankName,
		NoRekening:    transaction.NoRekening,
		Name:          transaction.Name,
		FundraisingID: transaction.FundraisingID,
		ImagePayment:  transaction.ImagePayment,
	}

	fundraising, err := s.adminRepo.FindFundraisingByID(int(transaction.FundraisingID))
	if err != nil {
		return transactionEntity, err
	}

	if transaction.Amount > fundraising.CurrentProgress {
		return transactionEntity, fmt.Errorf("the Amount Fundraising not enough")
	}

	newTransaction, err := s.transactionRepository.Save(transactionEntity)

	fundraising.CurrentProgress -= newTransaction.Amount
	_, err = s.adminRepo.UpdateFundraisingByID(fundraising.ID, fundraising)
	if err != nil {
		return transactionEntity, err
	}

	return newTransaction, err
}

func (s *transactionService) GetTransactionByID(id uint) (entities.Transaction, error) {
	transaction, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *transactionService) GetTransactions(limit int, offset int) ([]entities.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll(limit, offset)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
