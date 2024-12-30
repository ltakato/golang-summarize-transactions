package controllers

import (
	"github.com/gin-gonic/gin"
	"summarize-transactions/dto"
	"summarize-transactions/repositories"
)

type SummaryController struct {
	repository *repositories.TransactionsRepository
	BaseController
}

func NewSummaryController(repository *repositories.TransactionsRepository) *SummaryController {
	return &SummaryController{
		repository:     repository,
		BaseController: BaseController{},
	}
}

func (controller *SummaryController) GetSummary() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			availableDates, err := controller.repository.GetAvailableDates(c)

			response := dto.SummaryResponse{
				AvailableDates: availableDates,
			}

			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			controller.Ok(c, response)
		})
}

var _ IBaseController = (*SummaryController)(nil)
