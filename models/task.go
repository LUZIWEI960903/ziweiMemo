package models

import (
	"database/sql/driver"
	"time"
)

//type Time1 time.Time
//
//func (t Time1) MarshalText() (text []byte, err error) {
//	i := time.Time(t)
//	return []byte(i.Format("2006-01-02 15:04:05")), nil
//}
//
//func (t *Time1) UnmarshalText(text []byte) error {
//	text1 := string(text)
//	text1 = strings.Replace(text1, " ", "T", 1)
//	text1 += "Z"
//	t1, err := time.Parse("2006-01-02 15:04:05", string(text))
//	*t = Time1(t1)
//	return err
//}

type Time1 time.Time

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func (t *Time1) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = Time1(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = Time1(now)
	return
}

func (t Time1) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// 写入 mysql 时调用
func (t Time1) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 检出 mysql 时调用
//func (t *Time1) Scan(v interface{}) error {
//	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
//	tTime, _ := time.Parse(TimeFormat, v.(time.Time).String())
//	//tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
//	*t = Time1(tTime)
//	return nil
//}

// 用于 fmt.Println 和后续验证场景
func (t Time1) String() string {
	return time.Time(t).Format(TimeFormat)
}

type Task struct {
	TaskID     int64  `json:"task_id" db:"task_id"`                    // 备忘录id
	UserId     int64  `json:"user_id,string" db:"user_id"`             // 创建者id
	Status     int32  `json:"status" db:"status"`                      // 完成状态，默认0（未完成），1（完成）
	Title      string `json:"title" db:"title" binding:"required"`     // 备忘录标题
	Content    string `json:"content" db:"content" binding:"required"` // 备忘录内容
	StartTime  Time1  `json:"start_time" db:"start_time"`              // 备忘录开始时间
	EndTime    Time1  `json:"end_time" db:"end_time"`                  // 备忘录结束时间
	CreateTime Time1  `json:"create_time" db:"create_time"`            // 备忘录创建时间
	UpdateTime Time1  `json:"update_time" db:"update_time"`            // 备忘录最新修改时间
}

type UpdateTask struct {
	Status    int32  `json:"status,string" db:"status"`         // 完成状态，默认0（未完成），1（完成）
	Title     string `json:"title" db:"title"`                  // 备忘录标题
	Content   string `json:"content" db:"content"`              // 备忘录内容
	StartTime Time1  `json:"start_time,string" db:"start_time"` // 备忘录开始时间
	EndTime   Time1  `json:"end_time,string" db:"end_time"`     // 备忘录结束时间
}
