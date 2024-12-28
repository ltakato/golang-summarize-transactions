package dto

const UncategorizedCategoryToken = "(uncategorized)"

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
