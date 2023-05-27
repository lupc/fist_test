package test

import "fmt"

var i = initVar()

func init() {
	fmt.Println("init...")
}

func init() {
	fmt.Println("init2...")
}

func initVar() int {
	fmt.Println("initVar...")
	return 10
}

func RunInit() {
	fmt.Printf("i: %v\n", i)
	fmt.Println("RunInit...")
}
