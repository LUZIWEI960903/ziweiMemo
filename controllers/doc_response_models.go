package controllers

type _ResponseRegister struct {
	Code ResCode `json:"code" example:"1000"`     // 状态码
	Msg  string  `json:"msg" example:"Success!!"` // 信息
}

type _ResponseLogin struct {
	Code ResCode     `json:"code" example:"1000"`     // 状态码
	Msg  string      `json:"msg" example:"Success!!"` // 信息
	Data interface{} `json:"data"`                    // 数据
}

type _ResponseCreateTask struct {
	Code ResCode `json:"code" example:"1000"`     // 状态码
	Msg  string  `json:"msg" example:"Success!!"` // 信息
}
