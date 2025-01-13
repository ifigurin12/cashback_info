package card

import (
	api "cashback_info/app/rest_server/api/card"
	dto "cashback_info/interactor/dtos/card"
)

type ListUserCardsPresenter struct {
}

func (p *ListUserCardsPresenter) Present(outputDTO *dto.ListUserCardsOutputDTO) *api.ListUserCardsResoponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.ListUserCardsResoponse{
		Result: outputDTO.Items,
	}

	return response
}
