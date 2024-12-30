package controllers

import (
	"github.com/gin-gonic/gin"
	"summarize-transactions/dto"
	"summarize-transactions/repositories"
)

type CategoriesController struct {
	repository *repositories.TransactionsRepository
	BaseController
}

func NewCategoriesController(repository *repositories.TransactionsRepository) *CategoriesController {
	return &CategoriesController{
		repository:     repository,
		BaseController: BaseController{},
	}
}

func (controller *CategoriesController) GetCategories() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			categoryQuery := controller.categoryQuery(c)

			result, err := controller.repository.GetCategoriesWithTransactions(c, categoryQuery.Date)

			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			for i := range result {
				result[i].Normalize()
			}

			controller.Ok(c, result)
		})
}

func (controller *CategoriesController) GetCategoryTransactions() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			categoryId := c.Param("id")
			categoryQuery := controller.categoryQuery(c)

			result, err := controller.repository.GetCategoryTransactions(c, categoryId, categoryQuery.Date)

			if err != nil {
				controller.InternalServerError(c, nil)
				return
			}

			controller.Ok(c, result)
		})
}

func (controller *CategoriesController) categoryQuery(c *gin.Context) dto.CategoryQuery {
	categoryQuery, _ := c.Get("categoryQuery")
	return categoryQuery.(dto.CategoryQuery)
}

var _ IBaseController = (*CategoriesController)(nil)
