package main

import (
	"github.com/gin-gonic/gin"

)

func server(){
	gin.SetMode(gin.DebugMode)
	router:=gin.Default()

	Basic_REST(router)
	Advance_REST(router)
	group_func(router)
	test(router)
	router.Run(":9090")
}




func test(router *gin.Engine){

	router.POST("/bindJSON", funcBindJSON)

}



func main(){
	server()
}
