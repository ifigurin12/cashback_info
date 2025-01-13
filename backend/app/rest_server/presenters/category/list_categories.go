package category

import (
	api "cashback_info/app/rest_server/api/category"
	dto "cashback_info/interactor/dtos/category"
)

type ListCategoriesPresenter struct {
}

func (p *ListCategoriesPresenter) Present(outputDTO *dto.ListCategoriesOutputDTO) *api.ListCategoriesResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.ListCategoriesResponse{
		Result: outputDTO.Items,
	}

	return response
}
