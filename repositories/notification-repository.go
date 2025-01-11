package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"summarize-transactions/dto"
	"summarize-transactions/models"
)

type NotificationRepository struct {
	UserScopedRepository
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		UserScopedRepository: UserScopedRepository{
			db: db,
		},
	}
}

func (r *NotificationRepository) GetAll(c *gin.Context) ([]dto.NotificationResponse, error) {
	var err error
	var result []dto.NotificationResponse
	userId, _ := c.Get("userId")
	err = r.db.Model(&models.Notification{}).Select("id, text, read, created_at").Where(&models.Notification{UserID: userId.(string)}).Find(&result).Order("created_at desc").Error
	return result, err
}

func (r *NotificationRepository) CountUnread(c *gin.Context) (int64, error) {
	var err error
	var count int64
	userId, _ := c.Get("userId")
	err = r.db.Model(&models.Notification{}).Where(&models.Notification{UserID: userId.(string), Read: false}).Count(&count).Error
	return count, err
}
