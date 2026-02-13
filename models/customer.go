package models

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"size:100" json:"first_name"`
	LastName  string `gorm:"size:100" json:"last_name"`
	Email     string `gorm:"size:150;uniqueIndex" json:"email"`
	Phone     string `gorm:"size:20;uniqueIndex" json:"phone_number"`
}

