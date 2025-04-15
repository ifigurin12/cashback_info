package card

import (
	"cashback_info/internal/model/bank"
	"cashback_info/internal/model/user"
	"time"
)

type Card struct {
	ID            string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string    `gorm:"type:varchar(50);not null" json:"title"`
	UserID        string    `gorm:"type:uuid;not null" json:"user_id"`
	BankID        uint      `gorm:"not null" json:"bank_id"`
	DateCreated   time.Time `gorm:"type:timestamp;default:now();not null" json:"date_created"`
	LastUpdatedAt time.Time `gorm:"type:timestamp;default:now();not null" json:"last_updated_at"`
	User          user.User `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Bank          bank.Bank `gorm:"foreignKey:bank_id;constraint:OnDelete:CASCADE" json:"bank"`
}
