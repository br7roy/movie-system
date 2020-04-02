package model

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"msg" example:"请求信息"`
	Data    interface{} `json:"data"`
}
