package mysql

import (
	"database/sql"
	"ziweiMemo/models"
)

// CreateTask 插入新备忘录
func CreateTask(ts *models.Task) (err error) {
	sqlStr := `insert into task(task_id, title, user_id, content) values (?, ?, ?, ?);`

	_, err = db.Exec(sqlStr, ts.TaskID, ts.Title, ts.UserId, ts.Content)
	return
}

// ShowATaskByTaskID 根据taskid查询task信息
func ShowATaskByTaskID(taskId int64) (taskDetail *models.Task, err error) {
	taskDetail = new(models.Task)
	sqlStr := `select task_id, user_id, status, title, content, start_time, end_time, create_time, update_time from task where task_id = ?;`

	if err = db.Get(taskDetail, sqlStr, taskId); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorInvalidID
		}
	}
	return
}
