package models

const (
	OrderTime   = "time"
	OrderStatus = "status"
)

// RegisterParams 校验注册参数的结构体
type RegisterParams struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 再次确认密码
}

// LoginParams 校验登录参数的结构体
type LoginParams struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// TaskListParam 校验展示task参数的结构体
type TaskListParam struct {
	Page  int64  `json:"page"`  // 当前页数
	Size  int64  `json:"size"`  // 每页的展示数量
	Order string `json:"order"` // 排列顺序
}
