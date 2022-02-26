package stringutil

import (
	"fmt"
	"strconv"
	"strings"
)

func IntSlice2String(vals []int64, sep string) string {
	//var temp = make([]string, len(vals))
	if len(vals) == 0 {
		return ""
	}

	var temp string
	for _, v := range vals {
		temp += fmt.Sprintf("%s%d", sep, v)
	}

	temp += sep

	return temp
	//return strings.Join(temp, sep)
}

func StringSlice2String(vals []string, sep string) string {

	if len(vals) == 0 {
		return ""
	}

	var temp string
	for _, v := range vals {
		temp += fmt.Sprintf("%s%s", sep, v)
	}
	temp += sep
	return temp
}

func String2IntSlice(val string, sep string) []int64 {
	var temp []int64

	if len(val) > len(sep) {
		strSli := strings.Split(val, sep)

		for i := 0; i < len(strSli); i++ {
			if len(strSli[i]) == 0 {
				continue
			}
			intV, _ := strconv.ParseInt(strSli[i], 10, 64)
			temp = append(temp, intV)
		}
		//for _, v := range strSli {
		//	intV, _ := strconv.ParseInt(v, 10, 64)
		//	temp = append(temp, intV)
		//}
	}

	return temp
}

func String2StringSlice(val string, sep string) []string {
	var temp []string

	if len(val) > len(sep) {
		strSli := strings.Split(val, sep)

		for i := 0; i < len(strSli); i++ {
			if len(strSli[i]) == 0 {
				continue
			}
			temp = append(temp, strSli[i])
		}

		//for _, v := range strSli {
		//	intV, _ := strconv.ParseInt(v, 10, 64)
		//	temp = append(temp, intV)
		//}
	}

	return temp
}

func String_Int64(value string) int64 {
	valueInt, _ := strconv.ParseInt(value, 10, 64)
	return valueInt
}
