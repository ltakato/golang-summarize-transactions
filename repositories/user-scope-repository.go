package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"summarize-transactions/core"
)

type UserScopedRepository struct {
	db *gorm.DB
}

func (u *UserScopedRepository) userBaseParams(c *gin.Context) core.ParamsInterface {
	userId, _ := c.Get("userId")

	params := core.ParamsInterface{
		"userId": userId,
	}
	return params
}

func (u *UserScopedRepository) userScopedParams(c *gin.Context, addParams core.ParamsInterface) core.ParamsInterface {
	params := u.userBaseParams(c)
	for key, value := range addParams {
		params[key] = value
	}
	return params
}

func (u *UserScopedRepository) userScopedQuery(c *gin.Context, query string, params core.ParamsInterface) *gorm.DB {
	addedParams := u.userScopedParams(c, params)
	return u.db.Raw(query, addedParams)
}
