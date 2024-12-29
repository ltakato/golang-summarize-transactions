package dto

const UncategorizedCategoryToken = "(uncategorized)"

type SummaryResponse struct {
	AvailableDates []string `json:"availableDates"`
}

type CategoryResponse struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	TotalAmount float32 `json:"totalAmount"`
}

func (c *CategoryResponse) Normalize() *CategoryResponse {
	if c.ID == nil {
		uncategorized := UncategorizedCategoryToken
		c.ID = &uncategorized
		c.Name = &uncategorized
	}

	c.TotalAmount = c.TotalAmount / 100

	return c
}

type CategoryTransactionResponse struct {
	Title  string  `json:"title"`
	Date   string  `json:"date"`
	Amount float32 `json:"amount"`
}

func (c *CategoryTransactionResponse) Normalize() *CategoryTransactionResponse {
	c.Amount = c.Amount / 100

	return c
}
