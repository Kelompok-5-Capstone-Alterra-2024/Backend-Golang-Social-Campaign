package service

import (
	"capstone/entities"
	"capstone/repositories"
	"fmt"
)

type TransactionService interface {
	CreateTransaction(transaction entities.Transaction) (entities.Transaction, error)
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

func (s *transactionService) CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {

	fundraising, err := s.adminRepo.FindFundraisingByID(int(transaction.FundraisingID))
	if err != nil {
		return transaction, err
	}

	if transaction.Amount > fundraising.CurrentProgress {
		return transaction, fmt.Errorf("amount is greater than current progress")
	}

	newTransaction, err := s.transactionRepository.Save(transaction)

	fundraising.CurrentProgress -= newTransaction.Amount
	_, err = s.adminRepo.UpdateFundraisingByID(fundraising.ID, fundraising)
	if err != nil {
		return newTransaction, err
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
