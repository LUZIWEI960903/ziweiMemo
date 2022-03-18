package middleware

import (
	"strings"
	"time"
	"ziweiMemo/controllers"
	"ziweiMemo/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization: Bearer xxx.xxx.xxx
		authorHeader := c.Request.Header.Get("Authorization")
		if authorHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		tokenList := strings.SplitN(authorHeader, " ", 2)
		if !(len(tokenList) == 2 && tokenList[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 解析token
		claims, err := jwt.ParseToken(tokenList[1])
		if err != nil {
			if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
				controllers.ResponseError(c, controllers.CodeOverdueToken)
				c.Abort()
				return
			}
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.ContextUsernameKey, claims.Username)
		c.Set(controllers.ContextUserIDKey, claims.UserID)
		c.Next()
	}
}
