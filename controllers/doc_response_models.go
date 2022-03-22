package controllers

import "ziweiMemo/models"

type _ResponseSuccess struct {
	Code ResCode `json:"code" example:"1000"`     // 状态码
	Msg  string  `json:"msg" example:"Success!!"` // 信息
}

type _ResponseRegister struct {
	_ResponseSuccess
}

type _ResponseLogin struct {
	_ResponseSuccess
	Data interface{} `json:"data"` // 数据
}

type _ResponseCreateTask struct {
	_ResponseSuccess
}

type _ResponseShowATask struct {
	_ResponseSuccess
	Data *models.Task `json:"data"` // 具体task信息
}

type _ResponseShowAllTask struct {
	_ResponseSuccess
	Data []*models.Task `json:"data"` // 所有task信息
}

type _ResponseUpdateTask struct {
	_ResponseSuccess
}
