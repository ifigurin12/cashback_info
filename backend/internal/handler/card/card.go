package card

import (
	"cashback_info/internal/handler/card/api"
	entitycard "cashback_info/internal/model/card"
	cardservice "cashback_info/internal/service/card"
	cashbackservice "cashback_info/internal/service/category/cashback"
	familyservice "cashback_info/internal/service/family"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardHandler interface {
	SetupRoutes(router *gin.Engine)
}

type cardHandler struct {
	cardService             cardservice.CardService
	categoryCashbackService cashbackservice.CategoryCashbackService
	familyService           familyservice.FamilyService
}

func NewCardHandler(cardService cardservice.CardService, categoryCashbackService cashbackservice.CategoryCashbackService, familyService familyservice.FamilyService) CardHandler {
	return &cardHandler{cardService: cardService, categoryCashbackService: categoryCashbackService, familyService: familyService}
}

func (h *cardHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/cards", h.CreateCard)
	router.GET("/cards", h.ListCards)
	router.DELETE("/cards/:id", h.DeleteCard)
}

// CreateCard godoc
// @Summary Создает новую карту	пользователя
// @Description Создает новую карту
// @Security BearerAuth
// @Tags Card
// @Accept json
// @Produce json
// @Param request body api.CreateCardRequest true "Create Card Request"
// @Success 201 {object} api.CreateCardResponse
// @Router /cards [post]
func (h *cardHandler) CreateCard(c *gin.Context) {
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

	var request api.CreateCardRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardID, err := h.cardService.CreateCard(cardservice.CreateCardArgs{
		Title:  request.Title,
		UserID: userIDUUID,
		BankID: request.BankID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, api.CreateCardResponse{CardID: *cardID})
}

// ListUsers godoc
// @Summary Список карт пользователя
// @Description Вернет список кард, по идентификатору пользователя из токена
// @Security BearerAuth
// @Tags Card
// @Produce json
// @Success 200 {object} api.ListCardsResponse
// @Router /cards [get]
func (h *cardHandler) ListCards(c *gin.Context) {
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

	userCards, err := h.cardService.ListCards(userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	family, err := h.familyService.GetFamilyByUserID(userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var familyCards []entitycard.Card
	if family != nil {
		familyCards, err = h.cardService.ListCards(family.Leader.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, api.ListCardsResponse{
		UserCards:   userCards,
		FamilyCards: familyCards,
	})
}

// DeleteCard godoc
// @Summary Удалить карту по ID
// @Description Удаляет карту
// @Security BearerAuth
// @Tags Card
// @Accept json
// @Produce json
// @Param id path string true "Card ID"
// @Success 204
// @Router /cards/{id} [delete]
func (h *cardHandler) DeleteCard(c *gin.Context) {
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

	IsCardOwnedByUserID, err := h.cardService.IsCardOwnedByUserID(cardID, userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !*IsCardOwnedByUserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Card not owned by user"})
		return
	}

	err = h.cardService.DeleteCard(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.categoryCashbackService.Delete(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
