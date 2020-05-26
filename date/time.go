package date

import (
	"time"
)

// TimeMode time mode
type TimeMode string

const (
	Second TimeMode = "2006-01-02 15:04:05"
	Minute TimeMode = "2006-01-02 15:04"
	Hour   TimeMode = "2006-01-02 15"
	Day    TimeMode = "2006-01-02"
	Month  TimeMode = "2006-01"
	Year   TimeMode = "2006"
)

//TruncateTime Truncate time unix
func TruncateTime(ts int64, mode TimeMode) int64 {
	layout := string(mode)
	local, _ := time.LoadLocation("Asia/Shanghai")
	s := time.Unix(ts, 0).In(local).Format(layout)
	t, _ := time.ParseInLocation(layout, s, local)
	return t.Unix()
}
