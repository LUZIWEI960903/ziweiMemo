package controllers

import (
	"ziweiMemo/logic"
	"ziweiMemo/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CreateTaskHandler 创建task的接口
// @Summary 创建task的接口
// @Tags task接口
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param create_task body _RequestCreateTask true "需要上传的json"
// @Success 200 {object} _ResponseCreateTask
// @Router /task [post]
func CreateTaskHandler(c *gin.Context) {
	// 校验参数
	task := new(models.Task)
	if err := c.ShouldBindJSON(task); err != nil {
		zap.L().Error("[package: controllers] [func: CreateTaskHandler] [c.ShouldBindJSON(&task)] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userid, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	task.UserId = userid

	// 创建备忘录业务
	if err := logic.CreateTask(task); err != nil {
		zap.L().Error("[package: controllers] [func: CreateTaskHandler] [logic.CreateTask(task)] failed, ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, nil)
}
