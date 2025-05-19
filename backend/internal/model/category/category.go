package category

import (
	"cashback_info/internal/model/category/cashback"
	"cashback_info/internal/model/mcc"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Source      Source    `json:"source" binding:"required"`
	MCCCodes    []mcc.MCC `json:"mcc_codes" binding:"required"`
	Description *string   `json:"description,omitempty"`
}

type CategoryWithCashback struct {
	Category Category          `json:"category" binding:"required"`
	Cashback cashback.Cashback `json:"cashback" binding:"required"`
}

type Source string

const (
	SourceUser Source = "user"
	SourceBank Source = "bank"
)
