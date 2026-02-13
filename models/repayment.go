package models

import "time"

type Repayment struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID      uint      `gorm:"not null;index" json:"loan_id"`
	Loan        Loan      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Amount      float64   `gorm:"not null" json:"amount"`
	PaymentDate time.Time `gorm:"column:repayment_date;not null" json:"repayment_date"`
}

