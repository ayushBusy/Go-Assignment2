package services

import (
	"banking_system/models"

	"gorm.io/gorm"
)

type BankService struct {
	db *gorm.DB
}

func NewBankService(db *gorm.DB) *BankService {
	return &BankService{db: db}
}

func (s *BankService) Create(bank *models.Bank) error {
	return s.db.Create(bank).Error
}

func (s *BankService) GetByID(id uint) (*models.Bank, error) {
	var bank models.Bank
	if err := s.db.First(&bank, id).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func (s *BankService) GetAll() ([]models.Bank, error) {
	var banks []models.Bank
	if err := s.db.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}

func (s *BankService) Update(bank *models.Bank) error {
	return s.db.Save(bank).Error
}

func (s *BankService) Delete(id uint) error {
	return s.db.Delete(&models.Bank{}, id).Error
}

