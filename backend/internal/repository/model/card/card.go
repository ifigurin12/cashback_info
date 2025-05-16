package card

import (
	"cashback_info/internal/repository/model/bank"
	"time"
)

type Card struct {
	ID            string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string     `gorm:"type:varchar(50);not null" json:"title"`
	UserID        string     `gorm:"type:uuid;" json:"user_id"`
	Bank          *bank.Bank `gorm:"foreignKey:bank_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"bank"`
	DateCreated   time.Time  `gorm:"type:timestamp;default:now();not null" json:"date_created"`
	LastUpdatedAt time.Time  `gorm:"type:timestamp;default:now();not null" json:"last_updated_at"`
}

func (c *Card) TableName() string {
	return "cards"
}
