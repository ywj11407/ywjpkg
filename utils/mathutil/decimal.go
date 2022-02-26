package mathutil

import (
	"fmt"
	"strconv"
)

func FloatDecimal45(value float64, decimal int) float64 {

	temp := value + 0.0000000000001
	fmtStr := fmt.Sprintf("%%.%df", decimal)
	temp2, _ := strconv.ParseFloat(fmt.Sprintf(fmtStr, temp), 64)

	return temp2
}
