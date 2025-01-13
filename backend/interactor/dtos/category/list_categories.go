package category

import (
	entity "cashback_info/domain/entities/card"
)

type ListCategoriesInputDTO struct {
}

type ListCategoriesOutputDTO struct {
	Items []entity.Category `json:"categories"`
}
