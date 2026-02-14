package services

import (
	"errors"
	"time"

	"banking_system/models"

	"gorm.io/gorm"
)

type LoanService struct {
	db *gorm.DB
}

func NewLoanService(db *gorm.DB) *LoanService {
	return &LoanService{db: db}
}

func (s *LoanService) Create(loan *models.Loan) error {
	if loan.InterestRate == 0 {
		loan.InterestRate = 12.0
	}
	if loan.StartDate.IsZero() {
		loan.StartDate = time.Now()
	}
	if loan.Status == "" {
		loan.Status = "ongoing"
	}
	return s.db.Create(loan).Error
}

func (s *LoanService) GetByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	if err := s.db.First(&loan, id).Error; err != nil {
		return nil, err
	}
	return &loan, nil
}

func (s *LoanService) GetAll() ([]models.Loan, error) {
	var loans []models.Loan
	if err := s.db.Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

func (s *LoanService) GetCustomerLoans(customerID uint) ([]models.Loan, error) {
	var loans []models.Loan
	if err := s.db.Where("customer_id = ?", customerID).Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

func (s *LoanService) Update(loan *models.Loan) error {
	return s.db.Save(loan).Error
}

func (s *LoanService) Delete(id uint) error {
	return s.db.Delete(&models.Loan{}, id).Error
}

type LoanDetails struct {
	Loan               models.Loan `json:"loan"`
	TotalRepaid        float64     `json:"total_repaid"`
	LoanPending        float64     `json:"loan_pending"`
	InterestDueThisYear float64    `json:"interest_due_this_year"`
}

func (s *LoanService) GetDetails(id uint) (*LoanDetails, error) {
	loan, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	var totalRepaid float64
	if err := s.db.Model(&models.Repayment{}).
		Where("loan_id = ?", id).
		Select("COALESCE(SUM(amount), 0)").Scan(&totalRepaid).Error; err != nil {
		return nil, err
	}

	pending := loan.Amount - totalRepaid
	if pending < 0 {
		pending = 0
	}
	interestThisYear := pending * loan.InterestRate / 100.0

	return &LoanDetails{
		Loan:               *loan,
		TotalRepaid:        totalRepaid,
		LoanPending:        pending,
		InterestDueThisYear: interestThisYear,
	}, nil
}

func (s *LoanService) Repay(loanID uint, amount float64, paymentDate time.Time) (*models.Repayment, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	var repaymentRecord *models.Repayment

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var loan models.Loan
		if err := tx.First(&loan, loanID).Error; err != nil {
			return err
		}

		repayment := models.Repayment{
			LoanID:      loanID,
			Amount:      amount,
			PaymentDate: paymentDate,
		}
		if err := tx.Create(&repayment).Error; err != nil {
			return err
		}
		repaymentRecord = &repayment

		var totalRepaid float64
		if err := tx.Model(&models.Repayment{}).
			Where("loan_id = ?", loanID).
			Select("COALESCE(SUM(amount), 0)").Scan(&totalRepaid).Error; err != nil {
			return err
		}

		if totalRepaid >= loan.Amount && loan.Status != "closed" {
			loan.Status = "closed"
			if err := tx.Save(&loan).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return repaymentRecord, nil
}

