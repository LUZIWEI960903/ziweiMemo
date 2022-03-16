package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"ziweiMemo/models"
)

const secret = "https://github.com/LUZIWEI960903"

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

func InsertUser(user *models.User) (err error) {
	// 1. 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 2. sql语句
	sqlStr := `insert into user (user_id, username, password) values (?, ?, ?);`

	// 3. 插入数据
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// encryptPassword md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
