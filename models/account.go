package models

import "time"

type Account struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountNumber string    `gorm:"size:30;not null;uniqueIndex" json:"account_number"`
	BranchID      uint      `gorm:"not null;index" json:"branch_id"`
	Branch        Branch    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	AccountType   string    `gorm:"size:20;not null;default:savings" json:"account_type"`
	Interest      float64   `gorm:"not null;default:0" json:"interest"`
	Balance       float64   `gorm:"not null;default:0" json:"balance"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type AccountDetail struct {
	ID            uint `json:"account_id"`
	AccountNumber string `json:"account_number"`
	BranchID      uint `json:"branch_id"`
	AccountType   string `json:"account_type"`
	Interest      float64 `json:"interest"`
	Balance       float64 `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	Customers     []CustomerInfo `json:"customers"`
}

type CustomerInfo struct {
	CustomerID uint   `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone_number"`
	Role       string `json:"role"`
}

