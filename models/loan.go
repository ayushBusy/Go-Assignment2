package models

import "time"

type Loan struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID    uint      `gorm:"not null;index" json:"account_id"`
	Account      Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	CustomerID   uint      `gorm:"not null;index" json:"customer_id"`
	Customer     Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	Amount       float64   `gorm:"column:loan_amount;not null" json:"loan_amount"`
	InterestRate float64   `gorm:"column:loan_interest;not null" json:"loan_interest"` 
	StartDate    time.Time `gorm:"not null" json:"start_date"`
	TermMonths   int       `gorm:"not null" json:"term_months"`
	Status       string    `gorm:"size:20;not null" json:"status"`
}

