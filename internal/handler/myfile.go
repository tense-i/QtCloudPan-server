package handler

import (
	"QtCloudPan/internal/service"
	"QtCloudPan/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func ListHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	fmt.Println("ListHandler")
	// 打印请求携带数据
	path := r.URL.Path
	fmt.Println(path)
	// 拿到最后一个路径
	idx := strings.LastIndex(path, "/")
	if idx == -1 || idx == len(path)-1 {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid path")
		return
	}
	path = path[idx+1:]

	var req service.ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 调用服务层的列表逻辑
	response := service.ListFile(req, path)
	fmt.Println(response)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func ShareFileHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 GET 请求
	//if r.Method != http.MethodGet {
	//	utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only Get method is allowed")
	//	return
	//}

	fmt.Println("ShareFileHandler")

	var req service.ShareRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		fmt.Println(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	fmt.Println(req)

	// 调用服务层的分享逻辑
	response := service.ShareFile(req)
	fmt.Println(response)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
