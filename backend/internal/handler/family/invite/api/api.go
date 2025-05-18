package api

import "github.com/google/uuid"

type CreateFamilyInviteRequest struct {
	InviteeID uuid.UUID `json:"invitee_id" binding:"required"`
}
