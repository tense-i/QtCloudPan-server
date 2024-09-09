package main

import (
	"QtCloudPan/config"
	"QtCloudPan/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	// 加载配置文件
	config.LoadConfig()

	// 注册 HTTP 路由
	http.HandleFunc("/cloudObj/register", handler.RegisterHandler)
	http.HandleFunc("/cloudObj/login", handler.LoginHandler)
	http.HandleFunc("/cloudObj/myfiles/cound", handler.JWTMiddleware(handler.CoundHandler))

	// 启动服务器
	fmt.Println("Server is listening on port", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
