package family

import "cashback_info/internal/model/user"

type Family struct {
	ID      string      `json:"id" binding:"required"`
	Title   string      `json:"title" binding:"required"`
	Leader  user.User   `json:"leader" binding:"required"`
	Members []user.User `json:"members" binding:"required"`
}
