package service

import "QtCloudPan/internal/repository"

type CountRequest struct {
	Username string `json:"username"`
}

type CountResponse struct {
	Count int `json:"count"`
	Code  int `json:"code"`
}

// CountFile 统计文件数量
func CountFile(req CountRequest) CountResponse {
	// 调用数据访问层的统计逻辑
	countResp := repository.CountFile(req.Username)
	return CountResponse{
		Count: countResp.Count,
		Code:  countResp.Code,
	}
}
