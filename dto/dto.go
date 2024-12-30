package dto

import "summarize-transactions/core"

const UncategorizedCategoryToken = "(uncategorized)"

type SummaryResponse struct {
	AvailableDates []string `json:"availableDates"`
}

type CategoryResponse struct {
	ID          *string       `json:"id"`
	Name        *string       `json:"name"`
	TotalAmount core.Currency `json:"totalAmount"`
}

func (c *CategoryResponse) Normalize() *CategoryResponse {
	if c.ID == nil {
		uncategorized := UncategorizedCategoryToken
		c.ID = &uncategorized
		c.Name = &uncategorized
	}

	return c
}

type CategoryTransactionResponse struct {
	Title  string        `json:"title"`
	Date   string        `json:"date"`
	Amount core.Currency `json:"amount"`
}
