package api

import (
	"cashback_info/internal/model/category/cashback"

	"github.com/google/uuid"
)

type CreateCashbacksRequest struct {
	Cashbacks   []cashback.Cashback `json:"cashbacks" binding:"required"`
	CategoryIDs []uuid.UUID         `json:"category_ids" binding:"required"`
}

type UpdateCashbacksRequest struct {
	Cashbacks   []cashback.Cashback `json:"cashbacks" binding:"required"`
	CategoryIDs []uuid.UUID         `json:"category_ids" binding:"required"`
}
