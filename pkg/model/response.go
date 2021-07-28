package model

import "fmt"

// Response 接口响应
type Response interface {
	IsError() bool
	Error() string
}

// BaseResponse 通用返回格式
type BaseResponse struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

// IsError 判断接口是否返回错误
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// ErrorMsg 返回的错误信息
func (r BaseResponse) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}
