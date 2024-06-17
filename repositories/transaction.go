package repositories

import (
	"capstone/entities"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction entities.Transaction) (entities.Transaction, error)
	FindByID(id uint) (entities.Transaction, error)
	FindAll(limit int, offset int) ([]entities.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Save(transaction entities.Transaction) (entities.Transaction, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) FindByID(id uint) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.Where("id = ?", id).First(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionRepository) FindAll(limit int, offset int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	if err := r.db.Preload("Fundraising.Organization").Limit(limit).Offset(offset).Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}
