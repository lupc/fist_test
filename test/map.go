package test

import "fmt"

func RunMap() {
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("m: %v\n", m)

	fmt.Printf("m[\"a\"]: %v\n", m["a"])
	fmt.Printf("m[\"d\"]: %v\n", m["d"])
	v, ok := m["d"]
	if ok {
		fmt.Printf("v: %v\n", v)
	}
	v, ok = m["c"]
	if ok {
		fmt.Printf("v: %v\n", v)
	}
	fmt.Println("-----")
	for k, v := range m {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

}
