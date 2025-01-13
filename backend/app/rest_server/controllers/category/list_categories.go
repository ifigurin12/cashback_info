package card

import (
	utility "cashback_info/app/rest_server/controllers/private"
	presenter "cashback_info/app/rest_server/presenters/category"
	repocategory "cashback_info/infra/repos/category"
	dto "cashback_info/interactor/dtos/category"
	usecase "cashback_info/interactor/use_cases/category"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// @summary List Categories
// @Description Return categories
// @Tags Card
// @Accept json
// @Produce json
// @Success 200 {object} category.ListCategoriesResponse "Categories list"
// @Router /categories [get]
func (s *CategoryServer) ListCategories(c *gin.Context) {
	log.Info("CONTROLLER|ListCategories| Processing request to list categories")

	inputDTO := dto.ListCategoriesInputDTO{}
	presenter := &presenter.ListCategoriesPresenter{}

	categoryRepo := repocategory.NewCategoryRepo(s.postgresPool)

	useCase := usecase.ListCategoriesUseCase{
		CategoryRepo: categoryRepo,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)
	if err != nil {
		log.Error("CONTROLLER|ListCategories| Error while executing use case -> ", err)
		code, err := utility.TransformErrorToHttpError(err)
		c.AbortWithStatusJSON(code, gin.H{"error": err})
		return
	}

	response := presenter.Present(outputDTO)
	log.Info("CONTROLLER|ListCategories| Successfully retrieved categories")

	c.JSON(http.StatusOK, response)
}
