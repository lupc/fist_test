package test

import "fmt"

func RunBreakFor() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i: %v\n", i)
	}
}
