package controllers

import (
	"github.com/gin-gonic/gin"
	"summarize-transactions/core"
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
			tasks := []core.Task{
				func() (interface{}, error) {
					return controller.repository.GetAvailableDates(c)
				},
				func() (interface{}, error) {
					return controller.userRepository.GetUserInfo(c)
				},
			}
			results := core.RunConcurrentTasks(tasks)
			availableDates, err1 := results[0].Result, results[0].Error
			userInfo, err2 := results[1].Result, results[1].Error

			if err1 != nil || err2 != nil {
				controller.InternalServerError(c, nil)
				return
			}

			response := dto.SummaryResponse{
				User:           userInfo.(dto.UserInfo),
				AvailableDates: availableDates.([]string),
			}

			controller.Ok(c, response)
		})
}

var _ IBaseController = (*SummaryController)(nil)
