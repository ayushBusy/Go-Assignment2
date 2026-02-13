package services

import (
	"banking_system/models"

	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) Create(txn *models.Transaction) error {
	return s.db.Create(txn).Error
}

func (s *TransactionService) GetByID(id uint) (*models.Transaction, error) {
	var txn models.Transaction
	if err := s.db.First(&txn, id).Error; err != nil {
		return nil, err
	}
	return &txn, nil
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	var txns []models.Transaction
	if err := s.db.Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}

func (s *TransactionService) Update(txn *models.Transaction) error {
	return s.db.Save(txn).Error
}

func (s *TransactionService) Delete(id uint) error {
	return s.db.Delete(&models.Transaction{}, id).Error
}

