package logic

import (
	"ziweiMemo/dao/mysql"
	"ziweiMemo/models"
	"ziweiMemo/pkg/snowflake"
)

func UserRegister(p *models.RegisterParams) (err error) {
	// 1. 查询user是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 如果user不存在，则利用雪花算法生成唯一的userId
	userId := snowflake.GenID()
	// 3. 构造User实例
	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 4. 把User实例存入数据库并返回结果
	return mysql.InsertUser(user)
}
