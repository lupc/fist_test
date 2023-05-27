package test

import "fmt"

func RunClosure() {
	f1 := add(10)
	fmt.Printf("f1(1): %v\n", f1(1))
	fmt.Printf("f1(2): %v\n", f1(1))
	fmt.Printf("f1(2): %v\n", f1(1))

	fmt.Println("===============")
	f2 := add(20)
	fmt.Printf("f2(1): %v\n", f2(1))
	fmt.Printf("f2(1): %v\n", f2(1))

	fmt.Println("===============")
	fmt.Printf("f1(2): %v\n", f1(1))
	fmt.Printf("f1(2): %v\n", f1(1))
}

func add(a int) func(int) int {
	return func(b int) int {
		a += b
		return a
	}
}
