package controllers

type CategoryQuery struct {
	Date string `form:"date" binding:"required,partial_iso8601"`
}
