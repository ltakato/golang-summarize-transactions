package dto

import (
	"summarize-transactions/core"
)

const UncategorizedCategoryToken = "(uncategorized)"

type SummaryResponse struct {
	AvailableDates []string `json:"availableDates"`
}

type CategoryResponseText string

type CategoryResponse struct {
	ID          CategoryResponseText `json:"id"`
	Name        CategoryResponseText `json:"name"`
	TotalAmount core.Currency        `json:"totalAmount"`
}

func (c *CategoryResponseText) Normalize() CategoryResponseText {
	if *c == "" {
		*c = UncategorizedCategoryToken
	}
	return *c
}

func (c *CategoryResponse) Normalize() *CategoryResponse {
	c.ID.Normalize()
	c.Name.Normalize()
	return c
}

type CategoryTransactionResponse struct {
	Title  string        `json:"title"`
	Date   string        `json:"date"`
	Amount core.Currency `json:"amount"`
}
