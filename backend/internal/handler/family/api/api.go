package api

type CreateFamilyRequest struct {
	Title string `json:"title" binding:"required"`
}
