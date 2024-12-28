package repositories

import (
	"gorm.io/gorm"
	"summarize-transactions/dto"
)

type CategoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategoriesWithTransactions() ([]dto.CategoryResponse, error) {
	var err error
	query := `
			select
			c.id,
			c.name,
			sum(t.amount) total_amount
			from categories c
			right join transaction_categories tc on c.id  = tc.category_id
			right join transactions t on t.id = tc.transaction_id
			group by c.id
		`
	var result []dto.CategoryResponse
	err = r.db.Raw(query).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, err
}
