package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"ziweiMemo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createTaskHandler(c *gin.Context) {
	task := new(models.Task)
	if err := c.ShouldBindJSON(task); err != nil {
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

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	req.Header.Set(ContextUserIDKey, "936021537591296")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println(req.Header)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "936021537591296")
}

func showATaskHandler(c *gin.Context) {
	taskIdStr := c.Param("id")
	_, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		return
	}
	c.String(http.StatusOK, "参数解析成功")
}

func TestShowATaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/vi/task/:id"
	r.GET(url, showATaskHandler)

	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte(nil)))
	req.URL.Path = "/api/vi/task/1249050435260416"

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "参数解析成功")
}

func showAllTaskHandler(c *gin.Context) {
	p, err := getPageInfo(c)
	if err != nil {
		return
	}
	userIdStr := c.Request.Header.Get(ContextUserIDKey)
	c.JSON(http.StatusOK, gin.H{
		"userId": userIdStr,
		"page":   p.Page,
		"size":   p.Size,
	})
}

func TestShowAllTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/task"
	r.GET(url, showAllTaskHandler)

	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte(nil)))
	req.Header.Set(ContextUserIDKey, "453796320776192")
	req.URL.Query().Set("page", "1")
	req.URL.Query().Set("size", "2")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "\"userId\":\"453796320776192\"")
	assert.Contains(t, w.Body.String(), "\"page\":1")
	assert.Contains(t, w.Body.String(), "\"size\":2")
}

func updateTaskHandler(c *gin.Context) {
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		return
	}

	userIdStr := c.Request.Header.Get(ContextUserIDKey)

	updateTaskInfo := new(models.UpdateTask)
	if err := c.ShouldBindJSON(updateTaskInfo); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"taskId":         taskId,
		"userId":         userIdStr,
		"updateTaskInfo": updateTaskInfo,
	})
}

func TestUpdateTaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/task/:id"
	r.PUT(url, updateTaskHandler)

	body := `{"status":"1","title":"生活","content":"生活就像海洋","start_time":"2020-01-01 08:08:08","end_time":"2020-01-01 08:08:08"}`

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(body)))
	req.Header.Set(ContextUserIDKey, "936021537591296")
	req.URL.Path = "/api/v1/task/2389096492175360"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "2389096492175360")
	assert.Contains(t, w.Body.String(), "936021537591296")
	assert.Contains(t, w.Body.String(), body)
}

func deleteATaskHandler(c *gin.Context) {
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		return
	}

	userIdStr := c.Request.Header.Get(ContextUserIDKey)

	c.JSON(http.StatusOK, gin.H{
		"taskId": taskId,
		"userId": userIdStr,
	})
}

func TestDeleteATaskHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/task/:id"
	r.DELETE(url, deleteATaskHandler)

	req, _ := http.NewRequest(http.MethodDelete, url, bytes.NewReader([]byte(nil)))
	req.URL.Path = "/api/v1/task/2961887104864256"
	req.Header.Set(ContextUserIDKey, "453796320776192")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "2961887104864256")
	assert.Contains(t, w.Body.String(), "453796320776192")
}
