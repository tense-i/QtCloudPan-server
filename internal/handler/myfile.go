package handler

import (
	"QtCloudPan/internal/service"
	"QtCloudPan/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CoundHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	fmt.Println("CoundHandler")
	// 打印请求携带数据

	var req service.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 调用服务层的统计逻辑
	response := service.CountFile(req)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
