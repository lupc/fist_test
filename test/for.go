package test

import (
	"fmt"
	"time"
)

func RunFor() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
	}

}

// 相当于while
func RunFor2() {
	i := 1
	for i < 5 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

// 死循环
func RunFor3() {
	for {
		fmt.Println("死循环。。")
		time.Sleep(1 * time.Second)
	}
}

func RunForRange() {
	var arr = []byte{'a', 'b', 'c', 'd', 'e'}
	for i, v := range arr {
		fmt.Printf("i: %v,v:%c\n", i, v)
	}
}
