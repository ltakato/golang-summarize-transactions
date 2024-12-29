package repositories

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
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

func (r *CategoryRepository) GetCategoriesWithTransactions(date string) ([]dto.CategoryResponse, error) {
	var err error

	split := strings.Split(date, "-")
	year, err := strconv.ParseInt(split[0], 10, 32)
	month, err := strconv.ParseInt(split[1], 10, 32)

	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"year":  year,
		"month": month,
	}

	query := `
			select
			c.id,
			c.name,
			sum(t.amount) total_amount
			from categories c
			right join transaction_categories tc on c.id  = tc.category_id
			right join transactions t on t.id = tc.transaction_id
			where 
				EXTRACT(YEAR FROM t."date") = @year
  				AND EXTRACT(MONTH FROM t."date") = @month
			group by c.id
		`
	result := []dto.CategoryResponse{}
	err = r.db.Raw(query, params).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, err
}
