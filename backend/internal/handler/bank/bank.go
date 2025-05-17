package bank

import (
	bankservice "cashback_info/internal/service/bank"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankHandler interface {
	SetupRoutes(router *gin.Engine)
}

type bankHandler struct {
	bankService bankservice.BankService
}

func NewBankHandler(bankService bankservice.BankService) BankHandler {
	return &bankHandler{bankService: bankService}
}

func (b *bankHandler) SetupRoutes(router *gin.Engine) {
	router.GET("/banks", b.ListBanks)
}

// ListBanks обрабатывает запрос на получение списка банков
// @Summary Получение списка банков
// @Description Возвращает список банков
// @Tags Banks
// @Produce json
// @Success 200 {array} bank.Bank "Список банков"
// @Router /banks [get]
func (b *bankHandler) ListBanks(c *gin.Context) {
	banks, err := b.bankService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, banks)
}
