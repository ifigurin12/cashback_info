package category

import (
	entity "cashback_info/domain/entities/card"
)

type ListCategoriesResponse struct {
	Result []entity.Category `json:"result"`
}
