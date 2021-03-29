package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// uploadHandler用于处理上传文件请求
// 用于返回数据的Writer w，以及用于接受请求的指针r
func UploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
	// 	TODO: 返回上传html
		data, err := ioutil.ReadFile("static/view/upload.html")
		if err != nil {
			http.Error(w, "read file error", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, string(data))
	}else if r.Method == "POST" {
		// TODO: 接收文件流并存储到本地目录
	}
}
