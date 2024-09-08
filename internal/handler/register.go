package handler

import (
	service "QtCloudPan/internal/service"
	"QtCloudPan/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterHandler 处理用户注册的 HTTP 请求
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	fmt.Println("RegisterHandler")
	// 打印请求携带数据

	var req service.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 调用服务层的注册逻辑
	response := service.RegisterUser(req)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	fmt.Println("LoginHandler")
	// 打印请求携带数据

	var req service.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 调用服务层的登录逻辑
	response := service.LoginUser(req)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
