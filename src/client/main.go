package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"bytes"
	"mime/multipart"
	"os"
	"io"
)


// 用于读取resp的body
func helpRead(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(resp.Header)
	fmt.Println(string(body))
}

func upload(){
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw,_ := w.CreateFormFile("andrewfile", "upload_reboot.bat") //这里的uploadFile必须和服务器端的FormFile-name一致
	fd,_ := os.Open("/Users/eric/Desktop/reboot.bat")
	defer fd.Close()
	io.Copy(fw, fd)
	w.Close()
	resp,_ := http.Post("http://0.0.0.0:9090/upload", w.FormDataContentType(), buf)
	helpRead(resp)
}

func main(){
	// GET传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp,_ := http.Get("http://0.0.0.0:9090/getuser/TAO/101")
	helpRead(resp)

	// POST传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp,_ = http.Post("http://0.0.0.0:9090/form_post1?username=aiden&age=101", "",strings.NewReader(""))
	helpRead(resp)

	resp,_ = http.Post("http://0.0.0.0:9090/form_post2?username=aiden&age=101", "",strings.NewReader(""))
	helpRead(resp)

	resp,_ = http.Post("http://0.0.0.0:9090/bindJSON", "application/json", strings.NewReader("{\"user\":\"TAO\", \"password\": \"123\"}"))
	helpRead(resp)

	resp,_ = http.Post("http://0.0.0.0:9090/bindJSON", "application/json", strings.NewReader("{\"user\":\"TAO\", \"password\": \"123456\"}"))
	helpRead(resp)

	upload()
}


