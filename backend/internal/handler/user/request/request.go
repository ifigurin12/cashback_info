package request

type CreateUserRequest struct {
	Email    string  `json:"email" binding:"required"`
	Login    string  `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Phone    *string `json:"phone,omitempty"`
}
