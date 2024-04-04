package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册处理函数
	http.HandleFunc("/datastoreGet", postHandler_datastoreGet)
	http.HandleFunc("/keystoreSet", postHandler_keystoreSet)
	http.HandleFunc("/datastoreSet", postHandler_datastoreSet)
	http.HandleFunc("/keystoreGet", postHandler_keystoreGet)
	http.HandleFunc("/datastoreDelete", postHandler_datastoreDelete)

	// 启动 HTTP 服务器，监听在 8080 端口
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
