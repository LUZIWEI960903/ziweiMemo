package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	code resCode
	msg  interface{}
	data interface{}
}

// ResponseError 普通类型错误，默认msg
func ResponseError(c *gin.Context, code resCode) {
	c.JSON(http.StatusOK, &responseData{
		code: code,
		msg:  code.msg(),
		data: nil,
	})
}

// ResponseErrorWithMsg 自定义返回信息错误
func ResponseErrorWithMsg(c *gin.Context, code resCode, msg interface{}) {
	c.JSON(http.StatusOK, &responseData{
		code: code,
		msg:  msg,
		data: nil,
	})
}

// ResponseSuccess 成功，返回数据
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &responseData{
		code: codeSuccess,
		msg:  codeSuccess.msg(),
		data: data,
	})
}
