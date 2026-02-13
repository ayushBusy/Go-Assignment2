package models

import "time"

//AccountCustomer also handles the case for joint accounts where we specify the type of account (Current, Savings and Joint)
//basically the mapping b/w account and customer table
type AccountCustomer struct {
	AgreementID uint      `gorm:"primaryKey;autoIncrement" json:"agreement_id"`
	AccountID   uint      `gorm:"not null;index" json:"account_id"`
	Account     Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CustomerID  uint      `gorm:"not null;index" json:"customer_id"`
	Customer    Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Role        string    `gorm:"size:50;default:'primary_holder'" json:"role"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   string    `gorm:"size:100" json:"created_by"`
}

func (AccountCustomer) TableName() string {
	return "account_customers"
}

