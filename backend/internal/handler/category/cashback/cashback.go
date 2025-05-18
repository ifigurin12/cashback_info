package cashback

import (
	"cashback_info/internal/handler/category/cashback/api"
	cardservice "cashback_info/internal/service/card"
	cashbackservice "cashback_info/internal/service/category/cashback"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CashbackHandler interface {
	SetupRoutes(router *gin.Engine)
}

type cashbackHandler struct {
	cashbackCategoryService cashbackservice.CategoryCashbackService
	cardService             cardservice.CardService
}

func NewCashbackHandler(categoryService cashbackservice.CategoryCashbackService, cardService cardservice.CardService) CashbackHandler {
	return &cashbackHandler{cashbackCategoryService: categoryService, cardService: cardService}
}

func (h *cashbackHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/cards/:id/cashback", h.CreateCashbacks)
	router.PUT("/cards/:id/cashback", h.UpdateCashbacks)
}

// CreateCashbackCategories godoc
// @Summary Создание категорий для карты по ID
// @Description Возвращает 201 в случае успеха
// @Security BearerAuth
// @Tags cashback
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Param request body api.CreateCashbacksRequest true "Create Cashbacks Request"
// @Success 201
// @Router /cards/{id}/cashback [post]
func (h *cashbackHandler) CreateCashbacks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	userIDUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	cardIDStr := c.Param("id")
	cardID, err := uuid.Parse(cardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	var request api.CreateCashbacksRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Cashbacks) != len(request.CategoryIDs) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cashbacks and categories must have the same length"})
		return
	}

	isUserCard, err := h.cardService.IsCardOwnedByUserID(cardID, userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !*isUserCard {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not own the card"})
		return
	}

	err = h.cashbackCategoryService.Create(request.Cashbacks, request.CategoryIDs, cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

// UpdateCashbacks godoc
// @Summary Обновление категорий для карты по ID
// @Description Возвращает данные пользователя по указанному ID
// @Security BearerAuth
// @Tags cashback
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Param request body api.UpdateCashbacksRequest true "Create Cashbacks Request"
// @Success 204
// @Router /cards/{id}/cashback [put]
func (h *cashbackHandler) UpdateCashbacks(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	userIDUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID type"})
		return
	}

	cardIDStr := c.Param("id")
	cardID, err := uuid.Parse(cardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	var request api.CreateCashbacksRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(request.Cashbacks) != len(request.CategoryIDs) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cashbacks and categories must have the same length"})
		return
	}

	isUserCard, err := h.cardService.IsCardOwnedByUserID(cardID, userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !*isUserCard {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not own the card"})
		return
	}

	err = h.cashbackCategoryService.Delete(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.cashbackCategoryService.Create(request.Cashbacks, request.CategoryIDs, cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
