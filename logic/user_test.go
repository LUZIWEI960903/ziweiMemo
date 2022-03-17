package logic

import (
	"testing"
	"time"
	"ziweiMemo/models"

	sf "github.com/bwmarrin/snowflake"
)

var GenID int64

func init() {
	var st time.Time
	st, _ = time.Parse("2006-01-02", "2022-03-17")
	sf.Epoch = st.UnixNano() / 1000000
	node, _ := sf.NewNode(1)
	GenID = node.Generate().Int64()
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
