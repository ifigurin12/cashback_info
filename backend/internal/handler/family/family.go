package family

import (
	"cashback_info/internal/handler/family/api"
	familyservice "cashback_info/internal/service/family"
	familyuserservice "cashback_info/internal/service/family/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FamilyHandler interface {
	SetupRoutes(router *gin.Engine)
}

type familyHandler struct {
	familyService     familyservice.FamilyService
	familyUserService familyuserservice.FamilyUserService
}

func NewFamilyHandler(familyService familyservice.FamilyService, familyUserService familyuserservice.FamilyUserService) FamilyHandler {
	return &familyHandler{familyService: familyService, familyUserService: familyUserService}
}

func (f *familyHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/families", f.CreateFamily)
	router.GET("/families", f.GetFamily)
	router.DELETE("/families/:family-id", f.DeleteFamily)
	router.DELETE("/families/:family-id/members/:member-id", f.DeleteFamilyMember)
}

// CreateFamily godoc
// @Summary Создание семьи
// @Security BearerAuth
// @Description Создание семьи
// @Tags Family
// @Accept json
// @Produce json
// @Param request body api.CreateFamilyRequest true "Create family Request body"
// @Success 200 {object} family.Family
// @Router /families [post]
func (f *familyHandler) CreateFamily(c *gin.Context) {
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

	userFamily, err := f.familyService.GetFamilyByUserID(userIDUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userFamily != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already in family"})
		return
	}

	err = f.familyService.CreateFamily(request.Title, userIDUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetFamily godoc
// @Summary Получение семьи
// @Security BearerAuth
// @Description Получение семьи по id из Authorization header или family id
// @Tags Family
// @Accept json
// @Produce json
// @Param family-id query string false "Family ID"
// @Success 200 {object} family.Family
// @Router /families [get]
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

	familyIDStr := c.Query("family-id")
	if familyIDStr != "" {
		familyID, err := uuid.Parse(familyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
			return
		}

		family, err := h.familyService.GetFamilyByID(familyID)
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
		return
	} else {
		userFamily, err := h.familyService.GetFamilyByUserID(userIDUUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, userFamily)
		return
	}

}

// DeleteFamilyMember godoc
// @Summary Удаление члена семьи
// @Security BearerAuth
// @Description Удаление члена семьи по id
// @Tags Family
// @Accept json
// @Produce json
// @Param family-id path string true "Family ID"
// @Param member-id path string true "Member ID"
// @Success 204
// @Router /families/{family-id}/members/{member-id} [delete]
func (h *familyHandler) DeleteFamilyMember(c *gin.Context) {
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

	familyIDStr := c.Param("family-id")
	familyID, err := uuid.Parse(familyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
		return
	}

	memberIDStr := c.Param("member-id")
	memberID, err := uuid.Parse(memberIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid member ID"})
		return
	}

	family, err := h.familyService.GetFamilyByID(familyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if family.Leader.ID != userIDUUID {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not the leader"})
		return
	}

	err = h.familyUserService.Delete(familyID, memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteFamily godoc
// @Summary Удаление семьи
// @Security BearerAuth
// @Description Удаление семьи по id
// @Tags Family
// @Accept json
// @Produce json
// @Param family-id path string true "Family ID"
// @Success 204
// @Router /families/{family-id} [delete]
func (h *familyHandler) DeleteFamily(c *gin.Context) {
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

	familyIDStr := c.Param("family-id")
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

	isUserInFamily := family.IsUserInFamily(userIDUUID)
	if !isUserInFamily {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not in family"})
		return
	}

	err = h.familyService.DeleteFamily(familyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
