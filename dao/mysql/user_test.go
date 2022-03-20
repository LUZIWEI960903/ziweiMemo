package mysql

import (
	"testing"
	"ziweiMemo/models"
	"ziweiMemo/settings"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		Port:         ":3306",
		DbName:       "ziweiMemo",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	if err := Init(&dbCfg); err != nil {
		panic(err)
	}
}

func TestInsertUser(t *testing.T) {
	user := &models.User{
		UserID:   1,
		Username: "testName",
		Password: "123",
	}
	if err := InsertUser(user); err != nil {
		t.Fatalf("InsertUser(user) failed, err: %v\n", err)
	}
	t.Log("InsertUser(user) success!!")
}

func TestUserLogin(t *testing.T) {
	user := &models.User{
		UserID:   1,
		Username: "testName",
		Password: "123",
	}
	if err := UserLogin(user); err != nil {
		t.Fatalf("failed~~")
	}
	t.Log("success!!")
}
