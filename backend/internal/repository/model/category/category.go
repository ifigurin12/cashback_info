package category

import (
	"time"
)

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
