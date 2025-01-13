package card

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	BankType      BankType  `json:"bank_type"`
	DateCreated   time.Time `json:"date_created"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
}

type CardWithCategories struct {
	Card       Card       `json:"card"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	BankType    BankType  `json:"bank_type"`
	DateCreated time.Time `json:"date_created"`
	MCCCodes    []MССCode `json:"mcc_codes"`
	Description *string   `json:"description,omitempty"`
}

type CardCategory struct {
	CardID   uuid.UUID `json:"card_id"`
	Category Category  `json:"category"`
}

type MССCode struct {
	ID          int     `json:"id"`
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}
