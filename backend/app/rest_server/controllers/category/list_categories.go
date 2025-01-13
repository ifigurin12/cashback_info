package card

import (
	utility "cashback_info/app/rest_server/controllers/private"
	presenter "cashback_info/app/rest_server/presenters/category"
	repocategory "cashback_info/infra/repos/category"
	dto "cashback_info/interactor/dtos/category"
	usecase "cashback_info/interactor/use_cases/category"
	"net/http"

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

	inputDTO := dto.ListCategoriesInputDTO{}
	presenter := &presenter.ListCategoriesPresenter{}

	categoryRepo := repocategory.NewCategoryRepo(s.postgresPool)

	useCase := usecase.ListCategoriesUseCase{
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
