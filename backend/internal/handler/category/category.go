package category

import (
	categoryservice "cashback_info/internal/service/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	SetupRoutes(router *gin.Engine)
}

type categoryHandler struct {
	categoryService categoryservice.CategoryService
}

func NewCategoryHandler(categoryService categoryservice.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: categoryService}
}

func (h *categoryHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/categories", h.ListCategories)
}

// ListCategories обрабатывает получение списка категорий
// @Summary Получение списка категорий
// @Description Возвращает список категорий для указанного банка
// @Tags Category
// @Accept json
// @Produce json
// @Param bank-id query int true "Bank ID"
// @Success 200 {array} category.Category
// @Router /categories [get]
func (h *categoryHandler) ListCategories(c *gin.Context) {
	bankIDStr := c.Query("bank-id")
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank ID"})
		return
	}

	categories, err := h.categoryService.List(int32(bankID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
