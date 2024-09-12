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
		Code:  1,
	}
}

type ListResponse struct {
	List []Myfile `json:"list"`
	Code int      `json:"code"`
}

type Myfile struct {
	Username   string `json:"username"`
	Url        string `json:"url"`
	Size       int64  `json:"size"`
	FileName   string `json:"fileName"`
	Pv         int    `json:"pv"` // 下载量
	CreateTime string `json:"createTime"`
	Type       string `json:"type"`
}

// ListFile 获取文件列表
func ListFile(username string) ListResponse {
	// 在这里处理获取文件列表逻辑
	// 这里只是一个示例，实际上可能会更复杂
	return ListResponse{
		List: []Myfile{
			// 生成3个文件信息
			{
				Username:   username,
				Url:        "http://www.example.com/file1",
				Size:       1024,
				FileName:   "file1",
				Pv:         103,
				CreateTime: "2021-01-01 12:00:00",
				Type:       "txt",
			},
			{
				Username:   username,
				Url:        "http://www.example.com/file2",
				Size:       2048,
				FileName:   "file2",
				Pv:         200,
				CreateTime: "2021-01-02 12:00:00",
				Type:       "jpg",
			},
			{
				Username:   username,
				Url:        "http://www.example.com/file3",
				Size:       4096,
				FileName:   "file3",
				Pv:         1,
				CreateTime: "2021-01-03 12:00:00",
				Type:       "mp4",
			},
			{
				Username:   username,
				Url:        "http://www.example.com/file4",
				Size:       1,
				FileName:   "file4",
				Pv:         1000,
				CreateTime: "2021-01-04 12:00:00",
				Type:       "zip",
			},
		},
		Code: 1,
	}
}

type ShareResponse struct {
	Code        int `json:"code"`
	ShareStatus int `json:"shareStatus"`
}
type ShareRequest struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
	FileMd5  string `json:"filemd5"`
}

func ShareFile(username, filename, filemd5 string) ShareResponse {
	// 在这里处理分享文件逻辑

	// 这里只是一个示例，实际上可能会更复杂
	return ShareResponse{
		Code:        1,
		ShareStatus: 1,
	}
}
