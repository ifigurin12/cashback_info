package category

import (
	"time"

	"github.com/google/uuid"
)

type CategoryDB struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string     `gorm:"type:varchar(50);not null" json:"title"`
	BankID      *int32     `gorm:"type:int" json:"bank_id,omitempty"`
	UserID      *uuid.UUID `gorm:"type:uuid" json:"user_id,omitempty"`
	Description *string    `gorm:"type:text" json:"description,omitempty"`
	DateCreated time.Time  `gorm:"type:timestamp;default:now();not null" json:"date_created"`
}

func (CategoryDB) TableName() string {
	return "categories"
}
