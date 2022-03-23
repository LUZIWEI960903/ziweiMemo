package mysql

import (
	"database/sql"
	"fmt"
	"ziweiMemo/models"
)

// CreateTask 插入新备忘录
func CreateTask(ts *models.Task) (err error) {
	sqlStr := `insert into task(task_id, title, user_id, content) values (?, ?, ?, ?);`

	_, err = db.Exec(sqlStr, ts.TaskID, ts.Title, ts.UserId, ts.Content)
	return
}

// ShowATaskByTaskID 根据taskid查询task信息
func ShowATaskByTaskID(taskId, userId int64) (taskDetail *models.Task, err error) {
	taskDetail = new(models.Task)
	sqlStr := `select task_id, user_id, status, title, content, start_time, end_time, create_time, update_time from task where task_id = ? and is_deleted = 0;`

	if err = db.Get(taskDetail, sqlStr, taskId); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorInvalidID
		}
	}

	// 判断当前查询出来的task是否属于当前用户
	if taskDetail.UserId != userId {
		return nil, ErrorPermissionDenied
	}

	return
}

// GetTaskListByUserId 根据当前用户查询其所有task信息
func GetTaskListByUserId(userId int64, p *models.TaskListParam) (taskList []*models.Task, err error) {
	sqlStr := `select task_id, user_id, status, title, content, start_time, end_time, create_time, update_time from task where user_id = ? and is_deleted=0 order by create_time desc limit ?, ?;`

	taskList = make([]*models.Task, 0, p.Size)
	if err = db.Select(&taskList, sqlStr, userId, (p.Page-1)*p.Size, p.Size); err != nil {
		return nil, err
	}
	for _, v := range taskList {
		fmt.Println(*v)
	}

	return
}

// IsThisTaskBelongsUser 判断当前task是否属于当前用户
func IsThisTaskBelongsUser(taskId, userId int64) (err error) {
	task := new(models.Task)
	sqlStr := `select task_id, user_id, status, title, content, start_time, end_time, create_time, update_time from task where task_id = ? and is_deleted=0;`
	if err = db.Get(task, sqlStr, taskId); err != nil {
		if err == sql.ErrNoRows {
			return ErrorInvalidID
		}
	}
	fmt.Printf("%#v\n", task)
	// 判断当前查询出来的task是否属于当前用户
	if task.UserId != userId {
		return ErrorPermissionDenied
	}
	return nil
}

// UpdateTask 根据当前用户更新对应task的信息
func UpdateTask(taskId, userId int64, ts *models.UpdateTask) (err error) {
	// 查询该taskId
	if err = IsThisTaskBelongsUser(taskId, userId); err != nil {
		return err
	}

	// 更新数据
	sqlStr := `update task set status=?, title=?, content=?, start_time=?, end_time=? where task_id = ?;`
	_, err = db.Exec(sqlStr, ts.Status, ts.Title, ts.Content, ts.StartTime, ts.EndTime, taskId)

	fmt.Println(err)
	return
}

// DeleteATask 根据当前用户删除该条task
func DeleteATask(taskId, userId int64) (err error) {
	// 查询当前task是否为该用户
	if err = IsThisTaskBelongsUser(taskId, userId); err != nil {
		return err
	}

	// 更新该task的is_deleted字段
	sqlStr := `update task set is_deleted=1 where task_id = ?;`
	_, err = db.Exec(sqlStr, taskId)

	return
}
