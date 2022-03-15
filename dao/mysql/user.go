package mysql

// CheckUserExist 查询user是否存在
func CheckUserExist(username string) (err error) {
	// 1. sql语句
	sqlStr := `select count(*) from user where username = ?;`

	// 2. 查询
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return err
}
