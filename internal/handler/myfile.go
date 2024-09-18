package handler

import (
	"QtCloudPan/internal/Model"
	"QtCloudPan/internal/service"
	"QtCloudPan/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	// 调用服务层的分享逻辑
	response := service.ShareFile(req)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func DeleteFilesHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	fmt.Println("DeleteFilesHandler")

	var req service.DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 调用服务层的删除逻辑
	response := service.DeleteFiles(req)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 GET 请求
	if r.Method != http.MethodGet {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only Get method is allowed")
		return
	}

	fmt.Println("DownloadFileHandler")

	// 获取请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	filenames := r.Form.Get("filenames")

	// 调用服务层的下载逻辑
	response := service.DownloadFiles(username, filenames)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 请求
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}
	fmt.Println("UploadFileHandler")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取请求参数
	username := r.FormValue("username")
	uploadsDir := fmt.Sprintf("uploads/%s", username)
	if os.MkdirAll(uploadsDir, os.ModePerm) != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	files := r.MultipartForm.File["file"]
	var filesInfo []Model.Myfile
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		distPath := filepath.Join(uploadsDir, fileHeader.Filename)
		distFile, err := os.Create(distPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer distFile.Close()
		// 将文件流移动到开始位置
		if _, err := file.Seek(0, 0); err != nil {
			http.Error(w, "Error seeking file", http.StatusInternalServerError)
			return
		}
		if _, err := io.Copy(distFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileType := filepath.Ext(fileHeader.Filename)
		fileType = fileType[1:]
		filesInfo = append(filesInfo, Model.Myfile{
			Username: username,
			Url:      distPath,
			Size:     fileHeader.Size,
			FileName: fileHeader.Filename,
			Pv:       0,
			// 获取当前时间
			CreateTime: time.Now().String(),
			Type:       fileType,
		})

	}

	// 调用服务层的上传逻辑
	response := service.UploadFiles(filesInfo)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
