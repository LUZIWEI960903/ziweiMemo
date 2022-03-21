package controllers

import (
	"strconv"
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

// ShowATaskHandler 展示一条task详情的接口
// @Summary 展示一条task详情的接口
// @Tags task接口
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param   id     path    string     true        "id"
// @Success 200 {object} _ResponseShowATask
// @Router /task/{id} [get]
func ShowATaskHandler(c *gin.Context) {
	// 1. 解析参数
	taskIdStr := c.Param("id")
	taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	if err != nil {
		zap.L().Error("[package: controllers] [func: ShowATaskHandler] [c.Param(\"id\")] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 2. 业务逻辑
	taskData, err := logic.ShowATaskByTaskID(taskId, userId)
	if err != nil {
		zap.L().Error("[package: controllers] [func: ShowATaskHandler] [logic.ShowATaskHandler(taskId)] failed, ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, taskData)
}

// ShowAllTaskHandler 展示当前用户所有的task的接口
// @Summary 展示当前用户所有的task的接口
// @Tags task接口
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param page query string false "page"
// @Param size query string false "size"
// @Param order query string false "order"
// @Success 200 {object} _ResponseShowAllTask
// @Router /task [get]
func ShowAllTaskHandler(c *gin.Context) {
	// 1. 解析参数
	p, err := getPageInfo(c)
	if err != nil {
		zap.L().Error("[package: controllers] [func: ShowAllTaskHandler] [getPageInfo(c)] failed, ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 从JWT解析中获取userId
	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 3. 业务逻辑
	taskList, err := logic.GetTaskListByUserId(userId, p)
	if err != nil {
		zap.L().Error("[package: controllers] [func: ShowAllTaskHandler] [logic.GetTaskListByUserId(userId)] failed, ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 4. 返回响应
	ResponseSuccess(c, taskList)
}
