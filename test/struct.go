package test

import "fmt"

type WebSite struct {
	Name string //名称
	Url  string //地址
}

func RunStruct() {
	var site = WebSite{"百度", "http://www.baidu.com"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site) //%#v 详细输出 带结构体名称和属性名称
}
