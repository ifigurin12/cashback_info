package user

import (
	"cashback_info/internal/handler/user/request"
	userservice "cashback_info/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler interface {
	SetupRoutes(router *gin.Engine)
}

type userHandler struct {
	userService userservice.UserService
}

func NewUserHandler(userService userservice.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/users/:id", h.GetUserByID)
	router.POST("/users", h.CreateUser)
}

// GetUserByID обрабатывает получение пользователя по ID
// @Summary Получение пользователя по ID
// @Description Возвращает данные пользователя по указанному ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func (h *userHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	userID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser обрабатывает создание нового пользователя
// @Summary Создание нового пользователя
// @Description Создает нового пользователя и возвращает код
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.CreateUserRequest true "User data"
// @Success 201
// @Router /users [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	var request request.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.userService.CreateUser(
		request.Email,
		request.Login,
		request.Password,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
