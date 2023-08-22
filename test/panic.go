package test

import (
	"errors"
	"fmt"
)

func RunPanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic err: %v\n", err)
		}
	}()

	fmt.Println("运行。。。")
	panic(errors.New("程序异常终止！"))

}
