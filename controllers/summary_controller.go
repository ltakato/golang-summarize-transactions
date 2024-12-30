package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"summarize-transactions/dto"
	"summarize-transactions/repositories"
	"time"
)

func GetSummary(repository *repositories.TransactionsRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		availableDates, err := repository.GetAvailableDates(c)

		response := dto.SummaryResponse{
			AvailableDates: availableDates,
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
