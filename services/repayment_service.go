package services

import (
	"time"

	"banking_system/models"

	"gorm.io/gorm"
)

type RepaymentService struct {
	db *gorm.DB
}

func NewRepaymentService(db *gorm.DB) *RepaymentService {
	return &RepaymentService{db: db}
}

func (s *RepaymentService) Create(repayment *models.Repayment) error {
	if repayment.PaymentDate.IsZero() {
		repayment.PaymentDate = time.Now()
	}
	return s.db.Create(repayment).Error
}

func (s *RepaymentService) GetByID(id uint) (*models.Repayment, error) {
	var repayment models.Repayment
	if err := s.db.First(&repayment, id).Error; err != nil {
		return nil, err
	}
	return &repayment, nil
}

func (s *RepaymentService) GetAll() ([]models.Repayment, error) {
	var repayments []models.Repayment
	if err := s.db.Find(&repayments).Error; err != nil {
		return nil, err
	}
	return repayments, nil
}

func (s *RepaymentService) Update(repayment *models.Repayment) error {
	return s.db.Save(repayment).Error
}

func (s *RepaymentService) Delete(id uint) error {
	return s.db.Delete(&models.Repayment{}, id).Error
}

