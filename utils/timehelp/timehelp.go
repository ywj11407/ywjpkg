package timehelp

import (
	"fmt"
	"strconv"
	"time"
)

func GetDateTimeNowInt() int64 {
	StrDate := time.Now().Format("20060102150405")
	DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return DateNow
}

func GetDateIntFromDT(tm time.Time) int64 {
	StrDate := tm.Format("20060102")
	DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return DateNow
}

//20060102150405
func ChangeTimeInt2DT(tm int64) time.Time {
	localTime, _ := time.ParseInLocation("20060102150405", fmt.Sprintf("%d", tm), time.Local)
	return localTime
}

func ChangeDateTime2Int(tm time.Time) int64 {
	StrDate := tm.Format("20060102150405")
	DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return DateNow
}

//20060102150405
func FormatDateTime(dateTime int64) string {
	year := dateTime / 10000000000
	temp := dateTime % 10000000000
	month := temp / 100000000
	temp = temp % 100000000
	day := temp / 1000000
	temp = temp % 1000000
	hour := temp / 10000
	temp = temp % 10000
	minute := temp / 100
	second := temp % 100
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func GetDateTimeEndString() string {
	StrDate := time.Now().Format("20060102")
	StrDate += "000000"
	//DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return StrDate
}

func GetDateTimeStartInt() int64 {
	StrDate := time.Now().Format("20060102")
	StrDate += "000000"
	DateNow, _ := strconv.ParseInt(StrDate, 10, 64)
	return DateNow
}

func GetDateTimeMonthStartInt() string {
	strDay := time.Now().AddDate(0, -1, 0).Format("20060102")
	strDay += "000000"
	//DateNow, _ := strconv.ParseInt(strDay, 10, 64)
	//return DateNow
	return strDay
}

func GetDateTimeMonthEndInt() string {
	strDay := time.Now().Format("20060102")
	strDay += "235959"
	//DateNow, _ := strconv.ParseInt(strDay, 10, 64)
	//return DateNow
	return strDay
}

func GetDateNowInt() int64 {
	StrDate := time.Now().Format("20060102")
	DateNow, _ := strconv.Atoi(StrDate)
	return int64(DateNow)
}

func CalcTimeIntSpaceSecond(firstTime, secondTime int64) int64 {
	t1, _ := time.Parse("20060102150405", fmt.Sprintf("%d", firstTime))
	t2, _ := time.Parse("20060102150405", fmt.Sprintf("%d", secondTime))

	return t2.Unix() - t1.Unix()
}

func CalcTimeSpaceSecond(timeFormat string, firstTime, secondTime string) int64 {
	t1, _ := time.Parse(timeFormat, firstTime)
	t2, _ := time.Parse(timeFormat, secondTime)

	return t2.Unix() - t1.Unix()
}

func CalcTimeSpaceDay(timeFormat string, firstTime, secondTime string) int64 {
	t1, _ := time.Parse(timeFormat, firstTime)
	t2, _ := time.Parse(timeFormat, secondTime)
	day := int64(t2.Sub(t1).Hours() / 24)
	return day
}

func CalcTimeSpaceNatureDay(timeFormat string, firstTime, secondTime int64) int64 {

	//if firstTime/10000000000000000 > 1 {
	//	firstTime = firstTime / 1000
	//}
	//
	//if secondTime/10000000000000000 > 1 {
	//	secondTime = secondTime / 1000
	//}

	first := firstTime / 1000000
	first *= 1000000
	second := secondTime / 1000000
	second *= 1000000

	return CalcTimeSpaceDay(timeFormat, strconv.FormatInt(first, 10), strconv.FormatInt(second, 10))
}

func GetDateFromDateTime(dateTime int64) string {

	//if dateTime/10000000000000000 > 1 {
	//	dateTime = dateTime / 1000
	//}

	year := dateTime / 10000000000
	temp := dateTime % 10000000000
	month := temp / 100000000
	temp = temp % 100000000
	day := temp / 1000000
	//temp = temp % 1000000
	//hour := temp / 10000
	//temp = temp % 10000
	//minute := temp / 100
	//second := temp % 100
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func ChangeDateTimeToLimit(dateTime int64) int64 {
	dateTime = (dateTime / 1000000) * 1000000
	return dateTime + 235959
}

func GetFirstDateOfWeek(tm time.Time) time.Time {

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}

func GetLastDateOfWeek(tm time.Time) time.Time {

	weekStartDate := GetFirstDateOfWeek(tm)
	weekLastDate := time.Date(weekStartDate.Year(), weekStartDate.Month(), weekStartDate.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 6)
	return weekLastDate
}

func GetFirstDateOfMonth(tm time.Time) time.Time {

	monthStartDate := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local)

	return monthStartDate
}

func GetLastDateOfMonth(tm time.Time) time.Time {

	monthLastDate := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 1, 0).AddDate(0, 0, -1)

	return monthLastDate
}

func GetFirstDateOfQuarter(tm time.Time) time.Time {
	monthIndex := tm.Month()
	if monthIndex >= 1 && monthIndex <= 3 {
		monthIndex = 1
	} else if monthIndex >= 4 && monthIndex <= 6 {
		monthIndex = 4
	} else if monthIndex >= 7 && monthIndex <= 9 {
		monthIndex = 7
	} else {
		monthIndex = 10
	}

	quarterStartDate := time.Date(tm.Year(), monthIndex, 1, 0, 0, 0, 0, time.Local)
	return quarterStartDate
}

func GetLastDateOfQuarter(tm time.Time) time.Time {
	firstDate := GetFirstDateOfQuarter(tm)
	endDate := time.Date(firstDate.Year(), firstDate.Month(), 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 3, 0).AddDate(0, 0, -1)

	return endDate
}

func GetFirstDateOfYear(tm time.Time) time.Time {
	yearStartDate := time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return yearStartDate
}

func GetLastDateOfYear(tm time.Time) time.Time {
	yearLastDate := time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, time.Local).
		AddDate(1, 0, 0).AddDate(0, 0, -1)
	return yearLastDate
}

//*********************************
func GetFirstDateOfWeekInt(tmInt int64) int64 {

	tm := ChangeTimeInt2DT(tmInt)

	offset := int(time.Monday - tm.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	return ChangeDateTime2Int(weekStartDate)
}

func GetLastDateOfWeekInt(tmInt int64) int64 {
	tm := ChangeTimeInt2DT(tmInt)

	weekStartDate := GetFirstDateOfWeek(tm)
	weekLastDate := time.Date(weekStartDate.Year(), weekStartDate.Month(), weekStartDate.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, 6)
	return ChangeDateTime2Int(weekLastDate)
}

func GetFirstDateOfMonthInt(tmInt int64) int64 {

	tm := ChangeTimeInt2DT(tmInt)
	monthStartDate := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local)

	return ChangeDateTime2Int(monthStartDate)
}

func GetLastDateOfMonthInt(tmInt int64) int64 {

	tm := ChangeTimeInt2DT(tmInt)
	monthLastDate := time.Date(tm.Year(), tm.Month(), 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 1, 0).AddDate(0, 0, -1)

	return ChangeDateTime2Int(monthLastDate)
}

func GetFirstDateOfQuarterInt(tmInt int64) int64 {
	tm := ChangeTimeInt2DT(tmInt)
	monthIndex := tm.Month()
	if monthIndex >= 1 && monthIndex <= 3 {
		monthIndex = 1
	} else if monthIndex >= 4 && monthIndex <= 6 {
		monthIndex = 4
	} else if monthIndex >= 7 && monthIndex <= 9 {
		monthIndex = 7
	} else {
		monthIndex = 10
	}

	quarterStartDate := time.Date(tm.Year(), monthIndex, 1, 0, 0, 0, 0, time.Local)
	return ChangeDateTime2Int(quarterStartDate)
}

func GetLastDateOfQuarterInt(tmInt int64) int64 {
	tm := ChangeTimeInt2DT(tmInt)
	firstDate := GetFirstDateOfQuarter(tm)
	endDate := time.Date(firstDate.Year(), firstDate.Month(), 1, 0, 0, 0, 0, time.Local).
		AddDate(0, 3, 0).AddDate(0, 0, -1)

	return ChangeDateTime2Int(endDate)
}

func GetFirstDateOfYearInt(tmInt int64) int64 {
	tm := ChangeTimeInt2DT(tmInt)
	yearStartDate := time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return ChangeDateTime2Int(yearStartDate)
}

func GetLastDateOfYearInt(tmInt int64) int64 {
	tm := ChangeTimeInt2DT(tmInt)
	yearLastDate := time.Date(tm.Year(), 1, 1, 0, 0, 0, 0, time.Local).
		AddDate(1, 0, 0).AddDate(0, 0, -1)
	return ChangeDateTime2Int(yearLastDate)
}

//*****************************

func GetDateTimeInt64Year(tmInt int64) int64 {
	return tmInt / 10000000000
}

func GetDateTimeInt64Month(tmInt int64) int64 {
	temp := tmInt / 100000000
	temp = temp % 100
	return temp
}

func GetHourMinute(tm time.Time) int64 {
	now := time.Now()
	return int64(now.Hour()*100 + now.Minute())
}
