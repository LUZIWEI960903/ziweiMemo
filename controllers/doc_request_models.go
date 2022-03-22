package controllers

import "ziweiMemo/models"

type _RequestCreateTask struct {
	Title   string `json:"title" db:"title" binding:"required" example:"卷卷卷"`             // 备忘录标题
	Content string `json:"content" db:"content" binding:"required" example:"生命不惜，卷卷不止~~"` // 备忘录内容
}

type _RequestUpdateTask struct {
	Status    int32        `json:"status,string" db:"status"`         // 完成状态，默认0（未完成），1（完成）
	Title     string       `json:"title" db:"title"`                  // 备忘录标题
	Content   string       `json:"content" db:"content"`              // 备忘录内容
	StartTime models.Time1 `json:"start_time,string" db:"start_time"` // 备忘录开始时间
	EndTime   models.Time1 `json:"end_time,string" db:"end_time"`     // 备忘录结束时间
}
