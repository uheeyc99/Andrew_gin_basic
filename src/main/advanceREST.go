package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"log"
	"io"
)

func Advance_REST(router *gin.Engine){
	// 01  地址栏传递参数方法一   Param
	router.GET("getuser/:username/*age", func(context *gin.Context) {
		//http://127.0.0.1:9090/getuser/Aiden/5
		name:=context.Param("username")
		age:=context.Param("age")
		message:="hi " + name + " " + age
		fmt.Println(name)
		fmt.Println(age)
		context.String(http.StatusOK,message)
	})
	// 02  地址栏传递参数方法二    Query
	router.GET("getlocation", func(context *gin.Context) {
		Longitude:=context.DefaultQuery("lng","103.9") //默认值
		Latitude:=context.Query("lat")
		context.String(http.StatusOK,"your location is: %s,%s",Longitude,Latitude)
		//http://127.0.0.1:9090/getlocation?lng=103.8&lat=87.5
		//get post 通用
	})


	/* 03 body中传递数据 */
	router.POST("form_post1", func(context *gin.Context) {
		//"/form_post1?id=101&page=22"
		user:=context.PostForm("username")
		pass:=context.PostForm("passwd")
		//id:=context.Query("id")
		//page:=context.Query("page")
		//fmt.Println(user,pass,id,page)
		fmt.Println(user,pass)
		// post 回复 string
		context.String(http.StatusOK,"Okay "+user)
	})


	router.POST("form_post2", func(context *gin.Context) {
		user:=context.PostForm("username")
		pass:=context.PostForm("passwd")
		fmt.Println(user,pass)
		//context.JSON(http.StatusOK,gin.H{
		//	"token":"I am token",
		//	"message":"好自为之",
		//})


		type JsonHolder struct {
			Token string `json:"token" `
			Message string `json:"message"`
		}
		holder := JsonHolder{Token:"I am token!",Message:"好自为之啊"}
		context.JSON(http.StatusOK,holder)

	})

	/* 04 上传文件*/
	router.POST("upload", func(context *gin.Context) {
		//curl -X POST http://127.0.0.1:9090/upload -F "andrewfile=@/Users/eric/Desktop/reboot.bat" -H "Content-Type: multipart/form-data"
		file,fileHeader,err:=context.Request.FormFile("andrewfile")
		if err !=nil{
			fmt.Println("err1:",err)
			context.String(http.StatusBadRequest,"Bad resuest!")
			return
		}
		filename:=fileHeader.Filename
		fmt.Println("filename:",filename)

		out,err:=os.Create("upload"+"/"+filename)
		if err != nil{
			log.Fatal(err)
		}
		defer out.Close()
		_,err = io.Copy(out,file)
		if err != nil{
			log.Fatal(err)
		}
		context.String(http.StatusCreated,"upload successful")

		return
	})

}
