package family

import (
	"cashback_info/internal/handler/family/api"
	familyservice "cashback_info/internal/service/family"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FamilyHandler interface {
	SetupRoutes(router *gin.Engine)
}

type familyHandler struct {
	familyService familyservice.FamilyService
}

func NewFamilyHandler(familyService familyservice.FamilyService) FamilyHandler {
	return &familyHandler{familyService: familyService}
}

func (f *familyHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/families", f.CreateFamily)
	router.GET("/families/:id", f.GetFamily)
}

// CreateFamily godoc
// @Summary Создание семьи
// @Security BearerAuth
// @Description Создание семьи
// @Tags family
// @Accept json
// @Produce json
// @Param request body api.CreateFamilyRequest true "Create family Request body"
// @Success 200 {object} family.Family
// @Router /families [post]
func (h *familyHandler) CreateFamily(c *gin.Context) {
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

	var request api.CreateFamilyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.familyService.CreateFamily(request.Title, userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetFamily godoc
// @Summary Получение семьи
// @Security BearerAuth
// @Description Получение семьи по id
// @Tags family
// @Accept json
// @Produce json
// @Param id path string true "Family ID"
// @Success 200 {object} family.Family
// @Router /families/{id} [get]
func (h *familyHandler) GetFamily(c *gin.Context) {
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

	familyIDStr := c.Param("id")
	familyID, err := uuid.Parse(familyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	family, err := h.familyService.GetFamilyByID(familyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isUserInFamily := family.IsUserInFamily(userIDUUID)

	if !isUserInFamily {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not in family"})
		return
	}

	c.JSON(http.StatusOK, family)
}
