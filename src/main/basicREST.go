package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//Representational State Transfer
func Basic_REST(router *gin.Engine){
	router.GET("/get", func(context *gin.Context) {
		context.String(http.StatusOK,"get ok")
	})
	router.POST("/post", func(context *gin.Context) {
		context.String(http.StatusOK,"post OK")
	})
	router.PUT("/put", func(context *gin.Context) {
		context.String(http.StatusOK,"put OK")
	})
	router.DELETE("/delete", func(context *gin.Context) {
		context.String(http.StatusOK,"delete OK")
	})
	router.PATCH("/patch", func(context *gin.Context) {
		context.String(http.StatusOK,"patch OK")
	})
	router.HEAD("/head", func(context *gin.Context) {
		context.String(http.StatusOK,"head OK")
	})
	router.OPTIONS("/options", func(c *gin.Context) {
		c.String(http.StatusOK,"Options OK")

	})
}