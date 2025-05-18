package api

import (
	"github.com/google/uuid"
)

type CreateCardRequest struct {
	Title  string `json:"title" binding:"required"`
	BankID int32  `json:"bank_id" binding:"required"`
}

type CreateCardResponse struct {
	CardID uuid.UUID `json:"card_id" binding:"required"`
}
