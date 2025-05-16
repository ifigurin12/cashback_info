package card

import (
	"cashback_info/internal/model/bank"
	"cashback_info/internal/model/category"
	"time"
)

type Card struct {
	ID            string                          `json:"id" binding:"required"`
	Title         string                          `json:"title" binding:"required"`
	LastUpdatedAt time.Time                       `json:"last_updated_at" binding:"required"`
	Categories    []category.CategoryWithCashback `json:"categories" binding:"required"`
	Bank          *bank.Bank                      `json:"bank,omitempty"`
}
