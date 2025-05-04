package category

import (
	"time"
)

type CategoryCashback struct {
	CardID             string     `gorm:"type:uuid;not null" json:"card_id"`
	Category           Category   `gorm:"foreignKey:category_id;constraint:OnDelete:CASCADE" json:"category"`
	CashbackPercentage float64    `gorm:"type:decimal(5,1);not null;check:cashback_percentage > 0 AND cashback_percentage <= 100" json:"cashback_percentage"`
	StartDate          *time.Time `json:"start_date,omitempty"`
	EndDate            *time.Time `json:"end_date,omitempty"`
	CashbackLimit      *float64   `gorm:"type:int;check:cashback_limit > 0" json:"limit,omitempty"`
}

type Category struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `gorm:"type:varchar(50);not null" json:"title"`
	BankID      *uint     `gorm:"type:int" json:"bank_id,omitempty"`
	UserID      *string   `gorm:"type:uuid" json:"user_id,omitempty"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`
	DateCreated time.Time `gorm:"type:timestamp;default:now();not null" json:"date_created"`
}

func (Category) TableName() string {
	return "categories"
}

func (CategoryCashback) TableName() string {
	return "category_cashbacks"
}
