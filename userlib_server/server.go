package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// 创建一个文件，用于存储日志
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// 设置日志输出到文件
	log.SetOutput(logFile)

	// 注册处理函数
	http.HandleFunc("/datastoreGet", postHandler_datastoreGet)
	http.HandleFunc("/keystoreSet", postHandler_keystoreSet)
	http.HandleFunc("/datastoreSet", postHandler_datastoreSet)
	http.HandleFunc("/keystoreGet", postHandler_keystoreGet)
	http.HandleFunc("/datastoreDelete", postHandler_datastoreDelete)
	http.HandleFunc("/datastoreDeleteAll", postHandler_datastoreDeleteAll)
	http.HandleFunc("/keystoreDeleteAll", postHandler_keystoreDeleteAll)

	// 启动 HTTP 服务器，监听在 8080 端口
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
