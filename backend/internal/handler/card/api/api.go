package api

import (
	"cashback_info/internal/model/card"

	"github.com/google/uuid"
)

type CreateCardRequest struct {
	Title  string `json:"title" binding:"required"`
	BankID int32  `json:"bank_id" binding:"required"`
}

type CreateCardResponse struct {
	CardID uuid.UUID `json:"card_id" binding:"required"`
}

type ListCardsResponse struct {
	UserCards   []card.Card `json:"user_cards" binding:"required"`
	FamilyCards []card.Card `json:"family_cards" binding:"required"`
}
