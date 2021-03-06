package main
//http://www.okyes.me/2016/05/03/go-gin.html
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
)

// Binding数据
// 注意:后面的form:user表示在form中这个字段是user,不是User, 同样json:user也是
// 注意:binding:"required"要求这个字段在client端发送的时候必须存在,否则报错!
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
// bind JSON数据
func funcBindJSON(c *gin.Context) {
	var json Login
	// binding JSON,本质是将request中的Body中的数据按照JSON格式解析到json变量中
	if c.BindJSON(&json) == nil {
		if json.User == "TAO" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"JSON=== status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"JSON=== status": "unauthorized"})
		}
	} else {
		c.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
	}
}

// 下面测试bind FORM数据
func funcBindForm(c *gin.Context) {
	var form Login
	// 本质是将c中的request中的BODY数据解析到form中

	// 方法一: 对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
	// 方法二: 使用BindWith函数,如果你明确知道数据的类型
	if c.BindWith(&form, binding.Form) == nil{
		if form.User == "TAO" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"FORM=== status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"FORM=== status": "unauthorized"})
		}
	} else {
		c.JSON(404, gin.H{"FORM=== status": "binding FORM error!"})
	}
}

