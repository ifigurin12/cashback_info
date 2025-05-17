package invite

import (
	"cashback_info/internal/model/family"
	"cashback_info/internal/model/user"

	"github.com/google/uuid"
)

type FamilyInvite struct {
	ID     uuid.UUID     `json:"id" binding:"required"`
	Family family.Family `json:"family" binding:"required"`
	User   user.User     `json:"user" binding:"required"`
}
