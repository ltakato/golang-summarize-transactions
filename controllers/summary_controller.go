package controllers

import (
	"github.com/gin-gonic/gin"
	"summarize-transactions/dto"
	"summarize-transactions/repositories"
)

type SummaryController struct {
	userRepository *repositories.UserRepository
	repository     *repositories.TransactionsRepository
	BaseController
}

func NewSummaryController(userRepository *repositories.UserRepository, repository *repositories.TransactionsRepository) *SummaryController {
	return &SummaryController{
		userRepository: userRepository,
		repository:     repository,
		BaseController: BaseController{},
	}
}

func (controller *SummaryController) GetSummary() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			availableDates, err := controller.repository.GetAvailableDates(c)
			userInfo, err := controller.userRepository.GetUserInfo(c)

			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			response := dto.SummaryResponse{
				User:           userInfo,
				AvailableDates: availableDates,
			}

			controller.Ok(c, response)
		})
}

var _ IBaseController = (*SummaryController)(nil)
