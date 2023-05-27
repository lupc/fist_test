package test

import "fmt"

func RunArr() {

	var a1 = [...]int{2: 3, 4: 5}
	var a2 = [...]string{1: "b", 2: "c", 3: "d"}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}
