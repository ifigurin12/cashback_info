package cashback

import (
	"cashback_info/internal/repository/model/category"
	"time"

	"github.com/google/uuid"
)

type CategoryCashbackDB struct {
	CardID             uuid.UUID           `gorm:"type:uuid;not null" json:"card_id"`
	CategoryID         uuid.UUID           `gorm:"type:uuid;not null" json:"category_id"`
	Category           category.CategoryDB `gorm:"foreignKey:category_id;constraint:OnDelete:CASCADE" json:"category"`
	CashbackPercentage float64             `gorm:"type:decimal(5,1);not null;check:cashback_percentage > 0 AND cashback_percentage <= 100" json:"cashback_percentage"`
	StartDate          *time.Time          `json:"start_date,omitempty"`
	EndDate            *time.Time          `json:"end_date,omitempty"`
	CashbackLimit      *int32              `gorm:"type:int;check:cashback_limit > 0" json:"limit,omitempty"`
}

func (CategoryCashbackDB) TableName() string {
	return "category_cashbacks"
}
