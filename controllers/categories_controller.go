package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"summarize-transactions/repositories"
	"time"
)

func GetCategories(repository *repositories.TransactionsRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		categoryQuery, _ := c.Get("categoryQuery")

		result, err := repository.GetCategoriesWithTransactions(c, categoryQuery.(CategoryQuery).Date)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		for i := range result {
			result[i].Normalize()
		}

		c.JSON(http.StatusOK, result)
	}
}

func GetCategoryTransactions(repository *repositories.TransactionsRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		categoryId := c.Param("id")
		categoryQuery, _ := c.Get("categoryQuery")

		result, err := repository.GetCategoryTransactions(c, categoryId, categoryQuery.(CategoryQuery).Date)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
