package main

import (
	"fist_test/test"
	"fmt"
)

func main() {
	// fmt.Println("hello golang")
	// var msg = user.Hello()
	// fmt.Println(msg)
	// fmt.Printf("aaa.Aaa(): %v\n", aaa.Aaa())

	// test.RunIota2()
	// go test.RunInit()
	test.RunGoroutine1()

	//按q退出
	var q string
	for {
		fmt.Scanln(&q)
		if q == "q" {
			break
		}
	}

}
