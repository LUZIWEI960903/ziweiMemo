package logic

import (
	"testing"
	"time"
	"ziweiMemo/dao/mysql"
	"ziweiMemo/models"
	"ziweiMemo/settings"

	sf "github.com/bwmarrin/snowflake"
)

var GenID int64

func init() {
	var st time.Time
	st, _ = time.Parse("2006-01-02", "2022-03-17")
	sf.Epoch = st.UnixNano() / 1000000
	node, _ := sf.NewNode(1)
	GenID = node.Generate().Int64()

	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		Port:         ":3306",
		DbName:       "ziweiMemo",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	if err := mysql.Init(&dbCfg); err != nil {
		panic(err)
	}
}

func TestUserRegister(t *testing.T) {
	userId := GenID
	user := &models.User{
		UserID:   userId,
		Username: "testuser",
		Password: "testpassword",
	}
	//user := &models.User{
	//	UserID:   userId,
	//	Username: "testuser111",
	//	Password: "testpassword",
	//}
	if !(user.UserID == userId && user.Username == "testuser" && user.Password == "testpassword") {
		t.Fatal("models.User error!!")
	}
}

func TestUserLogin(t *testing.T) {
	user := &models.LoginParams{
		Username: "zs1",
		Password: "123",
	}

	_, err := UserLogin(user)
	if err == mysql.ErrorUserNotExist {
		t.Log("success!!")
	}
}
