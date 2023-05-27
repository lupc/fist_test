package test

import (
	"fmt"
	"time"
)

func RunSwitch() {
	var day = time.Now().Weekday()
	fmt.Printf("day type: %T\n", day)
	switch day {
	case time.Monday, time.Friday, time.Tuesday, time.Wednesday, time.Thursday:
		fmt.Println("工作日")
	case time.Sunday, time.Saturday:
		fmt.Println("休息日")
	}
}

// 带条件
func RunSwitchCond() {
	var score = 87
	switch {
	case score > 60:
		fmt.Println("及格")
		fallthrough //穿透到下一个case
	case score > 80 && score <= 90:
		fmt.Println("良好")
	case score > 90:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}
}
