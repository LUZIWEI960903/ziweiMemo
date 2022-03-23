package mysql

import (
	"testing"
	"ziweiMemo/models"
)

func TestCreateTask(t *testing.T) {
	task := &models.Task{
		TaskID:     123,
		UserId:     123,
		Status:     0,
		Title:      "just a test",
		Content:    "test test test~~",
		StartTime:  models.Time1{},
		EndTime:    models.Time1{},
		CreateTime: models.Time1{},
		UpdateTime: models.Time1{},
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

func TestUpdateTask(t *testing.T) {
	ts := &models.UpdateTask{
		Status:  0,
		Title:   "我是修改3333",
		Content: "我是修改的内容33333",
	}
	if err := UpdateTask(2687250877911040, 453796320776192, ts); err != nil {
		t.Fatal("failed~~")
	}
	t.Log("success!!")
}
