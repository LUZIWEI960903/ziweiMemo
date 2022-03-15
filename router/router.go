package router

import (
	"ziweiMemo/controllers"
	"ziweiMemo/logger"

	"github.com/gin-gonic/gin"
)

func SetUp(cfgMode string) *gin.Engine {
	if cfgMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// v1 版本
	v1 := r.Group("/api/v1")

	// 注册功能
	v1.POST("/register", controllers.UserRegisterHandler)

	return r
}
