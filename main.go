package main

import (
	"github.com/Winszheng/netdisk-go/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	// 监听所有网卡的8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("failed to start server: ", err)
	}
}
