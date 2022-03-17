package controllers

import (
	"errors"
	"fmt"
	"ziweiMemo/dao/mysql"
	"ziweiMemo/logic"
	"ziweiMemo/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandler 用户注册的接口
// @Summary 用户注册的接口
// @Description 如果用户存在则注册失败
// @Tags 用户注册的接口
// @Accept json
// @Produce json
// @Param register body models.RegisterParams true "需要上传的json"
// @Success 200 {object} _ResponseRegister
// @Router /register [post]
func UserRegisterHandler(c *gin.Context) {
	// 1. 解析参数
	p := new(models.RegisterParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserRegisterHandler] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	if err := logic.UserRegister(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserRegisterHandler] [logic.UserRegister(p)] failed, ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func UserLoginHandler(c *gin.Context) {
	// 1. 解析参数
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserLoginHandler] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	user, err := logic.UserLogin(p)
	if err != nil {
		zap.L().Error("[package: controllers] [func: UserLoginHandler] [logic.UserLogin] failed, ", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID), // int64 范围比前端的intiger 范围大 需要转换成string格式处理 不然会失真
		"username": user.Username,
		"token":    user.Token,
	})
}
