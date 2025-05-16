package mcc

type MCC struct {
	ID          int32   `json:"id" binding:"required"`
	Code        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}
