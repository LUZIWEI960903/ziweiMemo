package logic

import (
	"ziweiMemo/dao/mysql"
	"ziweiMemo/models"
	"ziweiMemo/pkg/jwt"
	"ziweiMemo/pkg/snowflake"
)

// UserRegister 处理user注册业务
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
	// 4. 把User实例传递到数据库处理并返回结果
	return mysql.InsertUser(user)
}

// UserLogin 处理user登录业务
func UserLogin(p *models.LoginParams) (user *models.User, err error) {
	// 1. 构造User实例
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 2. 把User实例传递到数据库处理
	if err := mysql.UserLogin(user); err != nil {
		return nil, err
	}

	// 3. 生成token
	token, err := jwt.GenToken(user.Username, user.UserID)

	if err != nil {
		return nil, err
	}
	// 4. 把生成的token赋值于user.Token中
	user.Token = token
	return
}

// ChangePassword 处理user修改密码业务
func ChangePassword(p *models.ChangePasswordParams) (err error) {
	// 1. 构造user实例
	user := &models.User{
		UserID:   p.UserId,
		Password: p.OPassword,
	}

	// 2. 判断旧密码是否正确
	if err = mysql.IsOpasswordCorrect(user); err != nil {
		return
	}

	// 3. 把新的密码更新进数据库
	return mysql.ChangePassword(p)
}
