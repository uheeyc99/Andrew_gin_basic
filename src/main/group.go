package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func Middleware(context *gin.Context){
	fmt.Println("this is middleware")
	name:=context.Query("name")
	if(name==""){
		context.String(http.StatusBadRequest,"which name ?")
		context.Abort()
	}
}
func group_func(router *gin.Engine){
	group01:=router.Group("/group1")
	group01.Use(Middleware)
	group01.GET("aaa", func(context *gin.Context) {
		name:=context.Query("name")
		fmt.Println(name)
		context.String(http.StatusOK,"name %s ok",name)
	})

}
