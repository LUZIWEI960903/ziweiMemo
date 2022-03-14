package router

import (
	"net/http"
	"ziweiMemo/logger"

	"github.com/gin-gonic/gin"
)

func SetUp(cfgMode string) *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, cfgMode)
	})
	return r
}
