package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
			result, err := controller.repository.GetAll(c)
			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			unreadCount, err := controller.repository.CountUnread(c)
			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			c.Header(XUnreadCount, fmt.Sprintf("%d", unreadCount))

			controller.Ok(c, result)
		})
}

var _ IBaseController = (*NotificationsController)(nil)
