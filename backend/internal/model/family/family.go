package family

import (
	"cashback_info/internal/model/user"

	"github.com/google/uuid"
)

type Family struct {
	ID     uuid.UUID `json:"id" binding:"required"`
	Title  string    `json:"title" binding:"required"`
	Leader user.User `json:"leader" binding:"required"`
}
