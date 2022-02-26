package timeutil

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateTimeMonday(t *testing.T) {

	fmt.Println(GetDateTimeMonday(time.Now()))
}

func TestChangeTimeInt2DT(t *testing.T) {
	tm := ChangeTimeInt2DT(20201101000000)
	fmt.Println("result:", tm.Format("20060102150405"))
}
