package controllers

import (
	"ziweiMemo/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	// 1. 解析参数
	p := new(models.RegisterParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[package: controllers] [func: UserRegisterHandler] [c.ShouldBindJSON(p)] failed, err: %v\n", zap.Error(err))
		ResponseError(c, codeInvalidParam)
	}
}
