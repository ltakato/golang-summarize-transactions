package dto

import (
	"summarize-transactions/core"
)

const UncategorizedCategoryToken = "(uncategorized)"

type UserInfo struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type ParseCsvRequest struct {
	CsvUrl string `json:"csvUrl"`
}

type SummaryResponse struct {
	User           UserInfo `json:"userInfo"`
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

type CategoryQuery struct {
	Date string `form:"date" binding:"required,partial_iso8601"`
}

type CategoryTransactionResponse struct {
	Title  string        `json:"title"`
	Date   string        `json:"date"`
	Amount core.Currency `json:"amount"`
}

type NotificationResponse struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
	Read      bool   `json:"read"`
}
