package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	ContextUsernameKey = "username"
	ContextUserIDKey   = "user_id"
)

var ErrorUserNeedLogin = errors.New("User need login!!")

// getCurrentUserID 获取当前用户id
func getCurrentUserID(c *gin.Context) (userid int64, err error) {
	useridStr, ok := c.Get(ContextUserIDKey)
	if !ok {
		return -1, ErrorUserNeedLogin
	}
	userid, ok = useridStr.(int64)
	if !ok {
		return -1, ErrorUserNeedLogin
	}
	return userid, nil
}
