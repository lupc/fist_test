package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	// 将当前时间转换为 Unix 时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	// 将指定时间转换为 Unix 时间戳
	t1, _ := time.Parse("2006-01-02 15:04:05", "2023-06-01 03:01:48")
	timestamp = t1.Unix()
	fmt.Println(timestamp)
}
