package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"summarize-transactions/repositories"
	"time"
)

func GetCategories(repository *repositories.CategoryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := repository.GetCategoriesWithTransactions()

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
