package card

import (
	entity "cashback_info/domain/entities/card"
)

type ListUserCardsResoponse struct {
	Result []entity.CardWithCategories `json:"result"`
}
