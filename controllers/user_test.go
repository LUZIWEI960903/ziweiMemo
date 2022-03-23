package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"ziweiMemo/models"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func userRegisterHandler(c *gin.Context) {
	p := new(models.RegisterParams)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	c.String(http.StatusOK, "参数校验成功")
}

func TestUserRegisterHandler(t *testing.T) {
	// 1. 构造一个接受请求的服务
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/register"
	r.POST(url, userRegisterHandler)

	// 2. 定义两个测试用例
	body := `{"username": "zs1", "password": "123", "re_Password": "123"}`

	// 3. 构造请求

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Equal(t, "参数校验成功", w.Body.String())
	assert.Contains(t, w.Body.String(), "参数校验成功")

}

func userLoginHandler(c *gin.Context) {
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		return
	}
	c.String(http.StatusOK, "参数校验成功")
}

func TestUserLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/login"
	r.POST(url, userLoginHandler)

	body := `{"username": "zs1", "password": "123"}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "参数校验成功")
}

func changePasswordHandler(c *gin.Context) {
	p := new(models.ChangePasswordParams)
	if err := c.ShouldBindJSON(p); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":     p.UserId,
		"o_password":  p.OPassword,
		"password":    p.Password,
		"re_password": p.RePassword,
	})
}

func TestChangePasswordHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/password"
	r.POST(url, changePasswordHandler)

	body := `{"user_id":"936021537591296","o_password":"123456","password":"12345","re_password":"12345"}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "936021537591296")
	assert.Contains(t, w.Body.String(), "123456")
	assert.Contains(t, w.Body.String(), "12345")
}
