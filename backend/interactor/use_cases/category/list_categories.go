package category

import (
	"context"

	dto "cashback_info/interactor/dtos/category"
	repo "cashback_info/interactor/ifaces/repos"

	log "github.com/sirupsen/logrus"
)

type ListCategoriesUseCase struct {
	CategoryRepo repo.ICategoryRepo
}

func (u *ListCategoriesUseCase) Execute(ctx context.Context, inputDTO dto.ListCategoriesInputDTO) (*dto.ListCategoriesOutputDTO, error) {
	result, err := u.CategoryRepo.ListCategories(ctx)

	if err != nil {
		log.Error("USE_CASE|ListUserCards|  Failed to list categories -> ", err)
		return nil, err
	}

	return &dto.ListCategoriesOutputDTO{
		Items: result,
	}, nil

}
