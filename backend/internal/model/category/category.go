package category

import (
	"cashback_info/internal/model/bank"
	"cashback_info/internal/model/card"
	"time"
)

type CategoryCashback struct {
	CardID             string    `gorm:"type:uuid;not null" json:"card_id"`
	CategoryID         string    `gorm:"type:uuid;not null" json:"category_id"`
	CashbackPercentage float64   `gorm:"type:decimal(5,1);not null;check:cashback_percentage > 0 AND cashback_percentage <= 100" json:"cashback_percentage"`
	StartDate          time.Time `json:"start_date,omitempty"`
	EndDate            time.Time `json:"end_date,omitempty"`
	Limit              float64   `gorm:"type:decimal(10,2)" json:"limit,omitempty"`
	Card               card.Card `gorm:"foreignKey:card_id;constraint:OnDelete:CASCADE" json:"card"`
	Category           Category  `gorm:"foreignKey:category_id;constraint:OnDelete:CASCADE" json:"category"`
}

type Category struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `gorm:"type:varchar(50);not null" json:"title"`
	BankID      uint      `gorm:"not null" json:"bank_id"`
	DateCreated time.Time `gorm:"type:timestamp;default:now();not null" json:"date_created"`
	Description string    `json:"description,omitempty"`
	Bank        bank.Bank `gorm:"foreignKey:bank_id;constraint:OnDelete:CASCADE" json:"bank"`
}
