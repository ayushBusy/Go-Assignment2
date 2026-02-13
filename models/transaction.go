package models

import "time"

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID   uint      `gorm:"not null;index" json:"account_id"`
	Account     Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	Type        string    `gorm:"size:20;not null" json:"transaction_type"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"column:transaction_date;autoCreateTime" json:"transaction_date"`
}

