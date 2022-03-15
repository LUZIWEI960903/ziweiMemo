package controllers

import (
	"errors"
	"ziweiMemo/dao/mysql"
	"ziweiMemo/logic"
	"ziweiMemo/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandler 用户登录的处理函数
func UserRegisterHandler(c *gin.Context) {
	// 1. 解析参数
	p := new(models.RegisterParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserRegisterHandler] [c.ShouldBindJSON(p)] failed, err: %v\n", zap.Error(err))
		ResponseError(c, codeInvalidParam)
	}

	// 2. 业务处理
	if err := logic.UserRegister(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserRegisterHandler] [login.UserRegister(p)] failed, err: %v\n", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {

		}
	}

	// 3. 返回响应
	ResponseSuccess(c, codeSuccess)
}
