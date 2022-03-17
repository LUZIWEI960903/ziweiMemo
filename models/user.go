package models

// User 用于存放数据库解析后的数据
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}
