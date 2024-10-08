package main

import (
	"QtCloudPan/config"
	"QtCloudPan/internal/handler"
	"QtCloudPan/internal/handler/middleware"
	"fmt"
	"net/http"
)

func main() {
	// 加载配置文件
	config.LoadConfig()
	// 注册 HTTP 路由
	http.HandleFunc("/cloudObj/register", handler.RegisterHandler)
	http.HandleFunc("/cloudObj/login", handler.LoginHandler)
	http.HandleFunc("/cloudObj/myfiles/count", middleware.JWTMiddleware(handler.CoundHandler))
	http.HandleFunc("/cloudObj/myfiles/list/", middleware.JWTMiddleware(handler.ListHandler))
	http.HandleFunc("/cloudObj/sharefile", middleware.JWTMiddleware(handler.ShareFileHandler))
	http.HandleFunc("/cloudObj/deletefiles", middleware.JWTMiddleware(handler.DeleteFilesHandler))
	http.HandleFunc("/cloudObj/downloadfiles", middleware.JWTMiddleware(handler.DownloadFileHandler))
	http.HandleFunc("/cloudObj/uploadfile", middleware.JWTMiddleware(handler.UploadFileHandler))

	// 启动服务器
	fmt.Println("Server is listening on port", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
