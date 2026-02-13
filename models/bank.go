package models

type Bank struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:100;not null;unique" json:"name"`
	Code     string `gorm:"size:20;unique" json:"code"`
	Location string `gorm:"size:120" json:"location"`
}

