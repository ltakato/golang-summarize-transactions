package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"summarize-transactions/core"
	"summarize-transactions/dto"
)

type UserRepository struct {
	UserScopedRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		UserScopedRepository: UserScopedRepository{
			db: db,
		},
	}
}

func (r *UserRepository) GetUserInfo(c *gin.Context) (dto.UserInfo, error) {
	var err error

	params := core.ParamsInterface{}
	query := `
		select id, email
		from users
		where id = @userId;	
	`
	var result dto.UserInfo
	err = r.userScopedQuery(c, query, params).Scan(&result).Error

	return result, err
}
