package card

import (
	utility "cashback_info/app/rest_server/controllers/private"
	"cashback_info/app/rest_server/presenters/card"
	repocard "cashback_info/infra/repos/card"
	repocategory "cashback_info/infra/repos/category"
	dto "cashback_info/interactor/dtos/card"
	usecase "cashback_info/interactor/use_cases/card"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @summary List User Cards
// @Description Return cards by user id
// @Tags Card
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} card.ListUserCardsResoponse "Cards list"
// @Router /users/{user_id}/cards [get]
func (s *CardServer) ListUserCards(c *gin.Context) {

	var userID uuid.UUID

	userIDParam := c.Param("user_id")
	userID, err := uuid.Parse(userIDParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	inputDTO := dto.ListUserCardsInputDTO{UserID: userID}
	presenter := &card.ListUserCardsPresenter{}

	cardRepo := repocard.NewCardRepo(s.postgresPool)
	categoryRepo := repocategory.NewCategoryRepo(s.postgresPool)

	useCase := usecase.ListUserCardsUseCase{
		CardRepo:     cardRepo,
		CategoryRepo: categoryRepo,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, err := utility.TransformErrorToHttpError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err})
		return
	}

	response := presenter.Present(outputDTO)

	c.JSON(http.StatusOK, response)
}
