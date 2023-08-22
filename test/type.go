package test

import "fmt"

func RunType() {
	var a = "hellow"
	var b = 1
	var c = true
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

type A struct {
	A1 string
}
type B struct {
	A
	B1 string
}

func RunTypeConvert() {

	var b = B{B1: "bbb"}
	b.A1 = "aaa"
	fmt.Printf("b.A: %v\n", b.A)

}
