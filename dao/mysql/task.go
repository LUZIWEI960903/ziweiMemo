package mysql

import (
	"ziweiMemo/models"
)

// CreateTask 插入新备忘录
func CreateTask(ts *models.Task) (err error) {
	sqlStr := `insert into task(task_id, title, user_id, content) values (?, ?, ?, ?);`

	_, err = db.Exec(sqlStr, ts.TaskID, ts.Title, ts.UserId, ts.Content)
	return
}
