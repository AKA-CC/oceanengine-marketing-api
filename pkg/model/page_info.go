package model

type PageInfo struct {
	// 页数
	Page int `json:"page,omitempty"`
	// 页面大小
	PageSize int `json:"page_size,omitempty"`
	// 总数
	TotalNumber int `json:"total_number,omitempty"`
	// 总页数
	TotalPage int `json:"total_page,omitempty"`
}
