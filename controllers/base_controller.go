package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BaseController struct{}

func (c *BaseController) Run(handler func(c *gin.Context)) gin.HandlerFunc {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return handler
}

func (c *BaseController) Ok(context *gin.Context, result interface{}) {
	context.JSON(http.StatusOK, result)
}

func (c *BaseController) InternalServerError(context *gin.Context, result interface{}) {
	context.JSON(http.StatusInternalServerError, result)
}
