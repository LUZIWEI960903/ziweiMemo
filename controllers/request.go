package controllers

import (
	"errors"
	"strconv"
	"ziweiMemo/models"

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

// getPageInfo 获取path上page，size信息
func getPageInfo(c *gin.Context) (*models.TaskListParam, error) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return nil, err
	}
	sizeStr := c.DefaultQuery("size", "2")
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return nil, err
	}

	p := &models.TaskListParam{
		Page:  page,
		Size:  size,
		Order: models.OrderTime,
	}
	return p, nil
}
