package service

import (
	"QtCloudPan/internal/repository"
	"fmt"
	"sort"
)

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

type ListRequest struct {
	Username string `json:"username"`
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

type ListResponse struct {
	List []Myfile `json:"list"`
	Code int      `json:"code"`
}

const (
	Asc  = "asc"
	Desc = "desc"
)

// ListFile 获取文件列表
func ListFile(req ListRequest, method string) ListResponse {
	// 调用数据访问层的列表逻辑
	listResp := repository.ListFile(req.Username)

	// 之后再优化
	var serviceFiles []Myfile
	for _, repoFile := range listResp.List {
		serviceFiles = append(serviceFiles, Myfile{
			Username:   repoFile.Username,
			Url:        repoFile.Url,
			Size:       repoFile.Size,
			FileName:   repoFile.FileName,
			Pv:         repoFile.Pv,
			CreateTime: repoFile.CreateTime,
			Type:       repoFile.Type,
		})
	}

	switch method {
	case Asc:
		// 升序
		sort.Slice(serviceFiles, func(i, j int) bool {
			return serviceFiles[i].Pv < serviceFiles[j].Pv
		})
	case Desc:
		// 降序
		sort.Slice(serviceFiles, func(i, j int) bool {
			return serviceFiles[i].Pv > serviceFiles[j].Pv
		})
	default:
		// 默认
	}
	fmt.Println(method)
	fmt.Println(serviceFiles)

	return ListResponse{
		List: serviceFiles,
		Code: listResp.Code,
	}
}

type ShareRequest struct {
	Username string `json:"username"`
	Filename string `json:"filename"`
	FileMd5  string `json:"filemd5"`
}
type ShareResponse struct {
	Code        int `json:"code"`
	ShareStatus int `json:"shareStatus"`
}

func ShareFile(req ShareRequest) ShareResponse {
	// 调用数据访问层的分享逻辑
	res := repository.ShareFile(req.Username, req.Filename, req.FileMd5)
	fmt.Println(res)
	return ShareResponse{
		Code:        res.Code,
		ShareStatus: res.ShareStatus,
	}
}
