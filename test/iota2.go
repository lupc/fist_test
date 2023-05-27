package test

import "fmt"

func RunIota2() {
	const (
		a1 = 1
		a2 = iota
		a3
		_
		a4
		a5 = iota + 1
		a6
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
}
