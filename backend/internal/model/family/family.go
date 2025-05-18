package family

import (
	"cashback_info/internal/model/user"

	"github.com/google/uuid"
)

type Family struct {
	ID      uuid.UUID   `json:"id" binding:"required"`
	Title   string      `json:"title" binding:"required"`
	Leader  user.User   `json:"leader" binding:"required"`
	Members []user.User `json:"members" binding:"required"`
}

func (f Family) IsUserInFamily(userID uuid.UUID) bool {
	if f.Leader.ID == userID {
		return true
	}

	for _, member := range f.Members {
		if member.ID == userID {
			return true
		}
	}
	return false
}
