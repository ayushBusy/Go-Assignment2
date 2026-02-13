package models

type Branch struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"size:100;not null" json:"branch_name"`
	Code    string `gorm:"size:20;unique" json:"code"`
	BankID  uint   `gorm:"not null;index" json:"bank_id"`
	Bank    Bank   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	Manager string `gorm:"size:120" json:"branch_manager"`
}

