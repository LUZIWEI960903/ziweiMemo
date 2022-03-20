package logic

import (
	"testing"
	"time"
	"ziweiMemo/models"

	"github.com/stretchr/testify/assert"

	sf "github.com/bwmarrin/snowflake"
)

func init() {
	var st time.Time
	st, _ = time.Parse("2006-01-02", "2022-03-17")
	sf.Epoch = st.UnixNano() / 1000000
	node, _ := sf.NewNode(1)
	GenID = node.Generate().Int64()
}

func TestCreateTask(t *testing.T) {
	taskId := GenID
	task := &models.Task{
		TaskID:     taskId,
		UserId:     936021537591296,
		Status:     0,
		Title:      "卷卷卷",
		Content:    "生命不惜，卷卷不止~~",
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
	assert.Equal(t, task.TaskID, GenID)
	assert.Equal(t, task.UserId, int64(936021537591296))
	assert.Equal(t, task.Status, int32(0))
	assert.Contains(t, task.Title, "卷")
	assert.Contains(t, task.Content, "卷")
}
