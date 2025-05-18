package invite

import (
	"cashback_info/internal/handler/family/invite/api"
	familyservice "cashback_info/internal/service/family"
	inviteservice "cashback_info/internal/service/family/invite"
	familyuserservice "cashback_info/internal/service/family/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FamilyInviteHandler interface {
	SetupRoutes(router *gin.Engine)
}

type familyInviteHandler struct {
	familyInviteService inviteservice.FamilyInviteService
	familyUserService   familyuserservice.FamilyUserService
	familyService       familyservice.FamilyService
}

func NewFamilyInviteHandler(familyInviteService inviteservice.FamilyInviteService, familyUserService familyuserservice.FamilyUserService, familyService familyservice.FamilyService) FamilyInviteHandler {
	return &familyInviteHandler{familyInviteService: familyInviteService, familyUserService: familyUserService, familyService: familyService}
}

func (f *familyInviteHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/families/:family-id/invites", f.CreateFamilyInvite)
	router.GET("/families/invites", f.GetFamilyInvite)
	router.DELETE("/families/:family-id/invites/:invite-id", f.DeleteFamilyInvite)
	router.POST("/families/:family-id/invites/accept", f.AcceptFamilyInvite)
	router.DELETE("/families/:family-id/invites/decline", f.DeclineFamilyInvite)

}

// CreateFamilyInvite godoc
// @Summary Создание приглашения в семью
// @Description Создание приглашения в семью
// @Security BearerAuth
// @Tags Family-Invite
// @Accept json
// @Produce json
// @Param request body api.CreateFamilyInviteRequest true "Request body"
// @Param family-id path string true "Family ID"
// @Success 201
// @Router /families/{family-id}/invites [post]
func (f *familyInviteHandler) CreateFamilyInvite(c *gin.Context) {
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

	var request api.CreateFamilyInviteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	family, err := f.familyService.GetFamilyByID(familyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
		return
	}
	if family.Leader.ID != userIDUUID {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not leader of family"})
		return
	}

	isUserInFamily, err := f.familyService.IsUserInFamily(request.InviteeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if *isUserInFamily {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already in family"})
		return
	}

	invites, err := f.familyInviteService.ListByFamilyID(familyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, invite := range invites {
		if invite.User.ID == request.InviteeID {
			c.JSON(http.StatusConflict, gin.H{"error": "User already invited"})
			return
		}
	}

	err = f.familyInviteService.Create(familyID, request.InviteeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetFamilyInvite godoc
// @Summary Получение приглашений
// @Description Получение приглашений либо по family-id либо по id из Authorization header
// @Security BearerAuth
// @Tags Family-Invite
// @Accept json
// @Produce json
// @Param family-id query string false "Family ID"
// @Success 204
// @Router /families/invites [get]
func (f *familyInviteHandler) GetFamilyInvite(c *gin.Context) {
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

		family, err := f.familyService.GetFamilyByID(familyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
			return
		}
		if family.Leader.ID != userIDUUID {
			c.JSON(http.StatusForbidden, gin.H{"error": "User is not leader of family"})
			return
		}

		familyInvites, err := f.familyInviteService.ListByFamilyID(familyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, familyInvites)
		return
	}

	userInvites, err := f.familyInviteService.ListByUserID(userIDUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userInvites)
}

// DeleteFamilyInvite godoc
// @Summary Удаление приглашения
// @Description Удаление приглашения, будет успешным только для лидера семьи
// @Security BearerAuth
// @Tags Family-Invite
// @Accept json
// @Produce json
// @Param family-id path string true "Family ID"
// @Param invite-id path string true "Invite ID"
// @Success 204
// @Router /families/{family-id}/invites/{invite-id} [delete]
func (f *familyInviteHandler) DeleteFamilyInvite(c *gin.Context) {
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

	inviteIDStr := c.Param("invite-id")
	inviteID, err := uuid.Parse(inviteIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invite ID"})
		return
	}

	family, err := f.familyService.GetFamilyByID(familyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid family ID"})
		return
	}
	if family.Leader.ID != userIDUUID {
		c.JSON(http.StatusForbidden, gin.H{"error": "User is not leader of family"})
		return
	}

	err = f.familyInviteService.DeleteByID(inviteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// AcceptFamilyInvite godoc
// @Summary Принятие приглашения
// @Description Принятие приглашения добавляет юзера в семью
// @Security BearerAuth
// @Tags Family-Invite
// @Accept json
// @Produce json
// @Param family-id path string true "Family ID"
// @Success 204
// @Router /families/{family-id}/invites/accept [post]
func (f *familyInviteHandler) AcceptFamilyInvite(c *gin.Context) {
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

	items, err := f.familyInviteService.ListByFamilyID(familyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inviteID *uuid.UUID

	for _, item := range items {
		if item.User.ID == userIDUUID {
			inviteID = &item.ID
			break
		}
	}
	if inviteID == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite not found"})
		return
	}

	err = f.familyInviteService.DeleteByID(*inviteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = f.familyUserService.Create(familyID, userIDUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// DeclineFamilyInvite godoc
// @Summary Отклонение приглашения
// @Description Удаляет приглашение
// @Security BearerAuth
// @Tags Family-Invite
// @Accept json
// @Produce json
// @Param family-id path string true "Family ID"
// @Success 204
// @Router /families/{family-id}/invites/decline [delete]
func (f *familyInviteHandler) DeclineFamilyInvite(c *gin.Context) {
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

	items, err := f.familyInviteService.ListByFamilyID(familyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inviteID *uuid.UUID

	for _, item := range items {
		if item.User.ID == userIDUUID {
			inviteID = &item.ID
			break
		}
	}

	if inviteID == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invite not found"})
		return
	}

	err = f.familyInviteService.DeleteByID(*inviteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
