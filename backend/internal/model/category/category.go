package category

import (
	"cashback_info/internal/model/category/cashback"
	"cashback_info/internal/model/mcc"
)

type Category struct {
	ID          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Source      Source    `json:"source" binding:"required"`
	MCCCodes    []mcc.MCC `json:"mcc_codes" binding:"required"`
	Description *string   `json:"description,omitempty"`
}

type CategoryWithCashback struct {
	Category
	cashback.Cashback
}

type Source string

const (
	SourceUser Source = "user"
	SourceBank Source = "bank"
)
