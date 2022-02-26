package timeutil

import (
	"strconv"
	"time"
)

func GetDateTimeNowInt() int64 {
	StrDate := time.Now().Format("20060102150405")
	DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return DateNow
}

func ChangeTimeInt2DT(tm int64) time.Time {
	year := tm / 10000000000
	temp := tm % 10000000000
	month := temp / 100000000
	temp = temp % 100000000
	day := temp / 1000000
	temp = temp % 1000000
	//hour := temp / 10000
	//temp = temp % 10000
	//minute := temp / 100
	//second := temp % 100

	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
}

func GetDateTimeMonday(tm time.Time) time.Time {

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//fmt.Println(weekStart)

	return weekStart
}

func GetDateTimeMonthFistDay(tm time.Time) time.Time {

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//fmt.Println(weekStart)

	return weekStart
}

func IsSameWeek(tm1, tm2 time.Time) bool {

	if GetDateTimeMonday(tm1) == GetDateTimeMonday(tm2) {
		return true
	}

	return false
}

func IsSameMonth(tm1, tm2 time.Time) bool {
	return true
}
