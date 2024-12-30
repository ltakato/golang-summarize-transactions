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
		userId, _ := c.Get("userId")

		defer cancel()

		categoryQuery, _ := c.Get("categoryQuery")

		result, err := repository.GetCategoriesWithTransactions(userId.(string), categoryQuery.(CategoryQuery).Date)

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
		categoryId := c.Param("id")
		userId, _ := c.Get("userId")
		categoryQuery, _ := c.Get("categoryQuery")

		defer cancel()

		result, err := repository.GetCategoryTransactions(userId.(string), categoryId, categoryQuery.(CategoryQuery).Date)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
