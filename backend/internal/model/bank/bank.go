package bank

type Bank struct {
	ID   int32  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
