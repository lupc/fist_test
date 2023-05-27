package user

import "fmt"

func Hello() string {
	var name = "小明"
	var age = 16
	return fmt.Sprintf("Hello, %s, %d", name, age)
}
