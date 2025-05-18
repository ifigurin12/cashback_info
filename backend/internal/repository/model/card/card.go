package card

import (
	"cashback_info/internal/repository/model/bank"
	"time"

	"github.com/google/uuid"
)

type CardDB struct {
	ID            uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string       `gorm:"type:varchar(50);not null" json:"title"`
	UserID        uuid.UUID    `gorm:"type:uuid;" json:"user_id"`
	BankID        *int32       `gorm:"type:int" json:"bank_id"`
	Bank          *bank.BankDB `gorm:"foreignKey:BankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"bank"`
	DateCreated   time.Time    `gorm:"type:timestamp;default:now();not null" json:"date_created"`
	LastUpdatedAt time.Time    `gorm:"type:timestamp;default:now();not null" json:"last_updated_at"`
}

func (c *CardDB) TableName() string {
	return "cards"
}
