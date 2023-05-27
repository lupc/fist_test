package test

import (
	"bytes"
	"fmt"
	"strings"
)

func RunString() {
	var s = `
	<html>
	<body>
	</body>
	</html>
	`

	fmt.Printf("s: %v\n", s)
}

func RunStringFormat() {
	var a = "aaa"
	var b = "bbb"
	var c = fmt.Sprintf("%s,%s", a, b)
	fmt.Printf("c: %v\n", c)
}

func RunStringJoin() {
	var a = "aaa"
	var b = "bbb"
	var c = strings.Join([]string{a, b}, "-")
	fmt.Printf("c: %v\n", c)
}

func RunStringBuilder() {
	var sb bytes.Buffer
	sb.WriteString("aaa")
	sb.WriteString("bbb")
	fmt.Printf("sb.String(): %v\n", sb.String())
}

// 字符串切片
func RunStringSlice() {
	var a = "hello world"
	var n = 4
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a[n]: %v,%c\n", a[n], a[n])
	fmt.Printf("a[1:n]: %v\n", a[1:n])
	fmt.Printf("a[2:]: %v\n", a[2:])
	fmt.Printf("a[:n]: %v\n", a[:n])
}
