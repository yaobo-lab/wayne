package dto

import "fmt"

var _ error = &ErrorResult{}

// Error implements the Error interface.
func (e *ErrorResult) Error() string {
	return fmt.Sprintf("code:%d,subCode:%d,msg:%s", e.Code, e.SubCode, e.Msg)
}

type ErrorResult struct {
	// http code
	Code int `json:"code"`
	// The custom code
	SubCode int    `json:"subCode"`
	Msg     string `json:"msg"`
}

// OpenAPI 通用 失败 返回接口
// swagger:response responseState
type Failure struct {
	// in: body
	// Required: true
	Body struct {
		ResponseBase
		Errors []string `json:"errors"`
	}
}

type ResponseBase struct {
	Code int `json:"code"`
}

// OpenAPI 通用 成功 返回接口
// swagger:response responseSuccess
type Success struct {
	// in: body
	// Required: true
	Body struct {
		ResponseBase
	}
}
