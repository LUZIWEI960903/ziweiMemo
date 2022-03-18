package logic

import (
	"ziweiMemo/dao/mysql"
	"ziweiMemo/models"
	"ziweiMemo/pkg/snowflake"
)

// CreateTask 创建备忘录的逻辑
func CreateTask(task *models.Task) (err error) {
	// 使用雪花算法生成taskid
	task.TaskID = snowflake.GenID()
	if err != nil {
		return err
	}
	// 去数据库插入数据
	if err := mysql.CreateTask(task); err != nil {
		return err
	}
	return
}
