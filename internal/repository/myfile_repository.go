package repository

type CountResponse struct {
	Count int `json:"count"`
	Code  int `json:"code"`
}

// CountFile 统计文件数量
func CountFile(username string) CountResponse {
	// 在这里处理统计文件数量逻辑
	// 这里只是一个示例，实际上可能会更复杂
	return CountResponse{
		Count: 100,
		Code:  200,
	}
}
