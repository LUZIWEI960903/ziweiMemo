package models

import "time"

type Task struct {
	TaskID     int64     `json:"task_id" db:"task_id"`                    // 备忘录id
	UserId     int64     `json:"user_id,string" db:"user_id"`             // 创建者id
	Status     int32     `json:"status" db:"status"`                      // 完成状态，默认0（未完成），1（完成）
	Title      string    `json:"title" db:"title" binding:"required"`     // 备忘录标题
	Content    string    `json:"content" db:"content" binding:"required"` // 备忘录内容
	StartTime  time.Time `json:"start_time" db:"start_time"`              // 备忘录开始时间
	EndTime    time.Time `json:"end_time" db:"end_time"`                  // 备忘录结束时间
	CreateTime time.Time `json:"create_time" db:"create_time"`            // 备忘录创建时间
	UpdateTime time.Time `json:"update_time" db:"update_time"`            // 备忘录最新修改时间
}
