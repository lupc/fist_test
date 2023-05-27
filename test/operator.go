package test

import "fmt"

func RunOperators() {
	a := 3
	b := 5

	var r = a & b
	var r2 = a | b
	var r3 = a ^ b

	var x = a << 2
	fmt.Printf("a: %b\n", a)
	fmt.Printf("b: %b\n", b)
	fmt.Printf("r: %b\n", r)
	fmt.Printf("r2: %b\n", r2)
	fmt.Printf("r3: %b\n", r3)
	fmt.Printf("x: %v\n", x)
}
