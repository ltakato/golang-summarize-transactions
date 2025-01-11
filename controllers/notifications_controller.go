package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"summarize-transactions/core"
	"summarize-transactions/repositories"
)

type NotificationsController struct {
	repository *repositories.NotificationRepository
	BaseController
}

const (
	XUnreadCount = "X-Unread-Count"
)

func NewNotificationsController(repository *repositories.NotificationRepository) *NotificationsController {
	return &NotificationsController{
		repository:     repository,
		BaseController: BaseController{},
	}
}

func (controller *NotificationsController) GetNotifications() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			tasks := []core.Task{
				func() (interface{}, error) {
					return controller.repository.GetAll(c)
				},
				func() (interface{}, error) {
					return controller.repository.CountUnread(c)
				},
			}
			results := core.RunConcurrentTasks(tasks)
			result1, err1 := results[0].Result, results[0].Error
			result2, err2 := results[1].Result, results[1].Error

			if err1 != nil || err2 != nil {
				controller.InternalServerError(c, "Failed to fetch notifications")
				return
			}

			c.Header(XUnreadCount, fmt.Sprintf("%d", result2))

			controller.Ok(c, result1)
		})
}

var _ IBaseController = (*NotificationsController)(nil)
