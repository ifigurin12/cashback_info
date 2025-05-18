package auth

import (
	"cashback_info/internal/handler/auth/request"
	authservice "cashback_info/internal/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SetupRoutes(router *gin.Engine)
}

type authHandler struct {
	authService authservice.AuthService
}

func NewAuthHandler(authService authservice.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) SetupRoutes(router *gin.Engine) {
	router.POST("/email/login", h.EmailAuth)
}

// EmailAuth обрабатывает аутентификацию пользователя по электронной почте
// @Summary Аутентификация пользователя по электронной почте
// @Description Возвращает токен и время истечения при успешной аутентификации
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.EmailLoginRequest true "Email and password"
// @Success 200 {object} user.Token
// @Router /email/login [post]
func (h *authHandler) EmailAuth(c *gin.Context) {
	var request request.EmailLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
