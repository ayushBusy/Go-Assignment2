package services

import (
	"banking_system/models"

	"gorm.io/gorm"
)

type BranchService struct {
	db *gorm.DB
}

func NewBranchService(db *gorm.DB) *BranchService {
	return &BranchService{db: db}
}

func (s *BranchService) Create(branch *models.Branch) error {
	return s.db.Create(branch).Error
}

func (s *BranchService) GetByID(id uint) (*models.Branch, error) {
	var branch models.Branch
	if err := s.db.First(&branch, id).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}

func (s *BranchService) GetAll() ([]models.Branch, error) {
	var branches []models.Branch
	if err := s.db.Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (s *BranchService) Update(branch *models.Branch) error {
	return s.db.Save(branch).Error
}

func (s *BranchService) Delete(id uint) error {
	return s.db.Delete(&models.Branch{}, id).Error
}

func (s *BranchService) GetByBankID(bankID uint) ([]models.Branch, error) {
	var branches []models.Branch
	if err := s.db.Where("bank_id = ?", bankID).Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}


