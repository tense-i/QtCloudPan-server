package repository

import (
	"QtCloudPan/internal/Model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

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

type DeleteResponse struct {
	Code         int `json:"code"`
	DeleteStatus int `json:"deleteStatus"`
}

func DeleteFiles(username string, filenames []string) DeleteResponse {
	// 在这里处理删除文件逻辑
	// 这里只是一个示例，实际上可能会更复杂
	return DeleteResponse{
		Code:         1,
		DeleteStatus: 1,
	}
}

type DownloadResponse struct {
	Code           int `json:"code"`
	DownloadStatus int `json:"downloadStatus"`
}

func DownloadFiles(username string, filenames string) DownloadResponse {
	// 在这里处理下载文件逻辑
	// 这里只是一个示例，实际上可能会更复杂
	return DownloadResponse{
		Code:           1,
		DownloadStatus: 1,
	}
}

type UploadResponse struct {
	Code         int `json:"code"`
	UploadStatus int `json:"uploadStatus"`
}

func UploadFiles(files []Model.Myfile) UploadResponse {
	// 将files写入到数据库

	for _, file := range files {
		var db *sql.DB
		db, err2 := sql.Open("mysql", "root:1352446@tcp(127.0.0.1:3306)/QtCloudPan")
		if err2 != nil {
			panic(err2)
		}
		_, err := db.Exec("INSERT INTO myfile (username, url, size, filename, pv, createTime, type) VALUES (?, ?, ?, ?, ?, ?, ?)", file.Username, file.Url, file.Size, file.FileName, file.Pv, file.CreateTime, file.Type)
		if err != nil {
			return UploadResponse{
				Code:         0,
				UploadStatus: 0,
			}
		}
	}

	return UploadResponse{
		Code:         1,
		UploadStatus: 1,
	}
}
