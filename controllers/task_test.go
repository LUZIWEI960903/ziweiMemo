package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"ziweiMemo/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func createTaskHandler(c *gin.Context) {
	task := new(models.Task)
	if err := c.ShouldBindJSON(task); err != nil {
		fmt.Println(task)
		zap.L().Error("[package: controllers] [func: CreateTaskHandler] [c.ShouldBindJSON(&task)] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userIdStr := c.Request.Header.Get(ContextUserIDKey)
	task.UserId, _ = strconv.ParseInt(userIdStr, 10, 64)

	fmt.Println(task)
	ResponseSuccess(c, task)
}

func TestCreateTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/task"
	r.POST(url, createTaskHandler)

	body := `{"title":"卷卷卷","content":"生命不惜，卷卷不止~~"}`

	req, _ := http.NewRequest("POST", url, bytes.NewReader([]byte(body)))
	req.Header.Set(ContextUserIDKey, "936021537591296")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(req.Header)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "936021537591296")
}
