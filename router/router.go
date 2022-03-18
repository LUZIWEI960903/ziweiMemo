package router

import (
	"ziweiMemo/controllers"
	"ziweiMemo/logger"
	"ziweiMemo/middleware"

	_ "ziweiMemo/docs" // 千万不要忘了导入把你上一步生成的docs

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func SetUp(cfgMode string) *gin.Engine {
	if cfgMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors())

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// v1 版本
	v1 := r.Group("/api/v1")

	// 注册功能
	v1.POST("/register", controllers.UserRegisterHandler)
	// 登录功能
	v1.POST("/login", controllers.UserLoginHandler)
	// 以下接口需要通过JWT认证后才能访问
	v1.Use(middleware.JWTAuthMiddleware())
	{
		// 创建task
		v1.POST("/task", controllers.CreateTaskHandler)
	}

	return r
}
