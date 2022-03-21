package mysql

import (
	"testing"
	"time"
	"ziweiMemo/models"
)

func TestCreateTask(t *testing.T) {
	task := &models.Task{
		TaskID:     123,
		UserId:     123,
		Status:     0,
		Title:      "just a test",
		Content:    "test test test~~",
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	}
	if err := CreateTask(task); err != nil {
		t.Fatal("failed~~")
	}
	t.Log("testing success!!")
}

func TestShowATaskByTaskID(t *testing.T) {
	taskDetail, err := ShowATaskByTaskID(1249050435260416, 437364308578304)
	if err != nil {
		t.Fatal("failed~~")
	}
	t.Log("success!!", taskDetail)
}

func TestGetTaskListByUserId(t *testing.T) {
	taskList, err := GetTaskListByUserId(437364308578304, &models.TaskListParam{
		Page:  1,
		Size:  2,
		Order: "time",
	})
	if err != nil {
		t.Fatal("failed~~")
	}
	for _, v := range taskList {
		t.Logf("%v\n", *v)
	}
}
