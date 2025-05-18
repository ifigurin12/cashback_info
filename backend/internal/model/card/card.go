package card

import (
	"cashback_info/internal/model/bank"
	"cashback_info/internal/model/category"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID            uuid.UUID                       `json:"id" binding:"required"`
	Title         string                          `json:"title" binding:"required"`
	LastUpdatedAt time.Time                       `json:"last_updated_at" binding:"required"`
	Categories    []category.CategoryWithCashback `json:"categories" binding:"required"`
	Bank          *bank.Bank                      `json:"bank,omitempty"`
}
