package category

import (
	entity "cashback_info/internal/model/category"
	repository "cashback_info/internal/repository/category"
	repositoryentity "cashback_info/internal/repository/model/category"

	"github.com/google/uuid"
)

type CategoryService interface {
	List(bankID int32) ([]entity.Category, error)
	Create(categories []entity.Category, userID uuid.UUID) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (c *categoryService) List(bankID int32) ([]entity.Category, error) {
	items, err := c.categoryRepository.List(bankID)
	if err != nil {
		return nil, err
	}

	result := make([]entity.Category, len(items))
	for i, item := range items {
		source := entity.SourceBank

		result[i] = entity.Category{
			ID:          item.ID,
			Title:       item.Title,
			Source:      source,
			Description: item.Description,
		}
	}

	return result, nil
}

func (c *categoryService) Create(categories []entity.Category, userID uuid.UUID) error {
	items := make([]repositoryentity.CategoryDB, len(categories))

	for i, category := range categories {
		items[i] = repositoryentity.CategoryDB{
			ID:          category.ID,
			Title:       category.Title,
			UserID:      &userID,
			Description: category.Description,
		}
	}

	return c.categoryRepository.Create(items)
}
