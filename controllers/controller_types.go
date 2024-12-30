package controllers

import "github.com/gin-gonic/gin"

type IBaseController interface {
	Run(handler func(c *gin.Context)) gin.HandlerFunc
	Ok(context *gin.Context, result interface{})
	InternalServerError(context *gin.Context, result interface{})
}
