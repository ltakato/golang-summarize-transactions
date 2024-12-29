package repositories

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
	"summarize-transactions/dto"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{
		db: db,
	}
}

func (r *TransactionsRepository) GetAvailableDates() ([]string, error) {
	var err error

	query := `
		select
			concat(EXTRACT(YEAR FROM t."date"), '-', EXTRACT(MONTH FROM t."date")) concat_date
		from transactions t
		group by concat_date;
	`
	var result []string
	err = r.db.Raw(query).Scan(&result).Error

	return result, err
}

func (r *TransactionsRepository) GetCategoriesWithTransactions(date string) ([]dto.CategoryResponse, error) {
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
