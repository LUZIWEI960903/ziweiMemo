package controllers

type _RequestCreateTask struct {
	UserId  int64  `json:"user_id,string" db:"user_id" binding:"required" example:"437364308578304"` // 创建者id
	Title   string `json:"title" db:"title" binding:"required" example:"卷卷卷"`                        // 备忘录标题
	Content string `json:"content" db:"content" binding:"required" example:"生命不惜，卷卷不止~~"`            // 备忘录内容
}
