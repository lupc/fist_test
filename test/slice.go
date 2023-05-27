package test

import "fmt"

func RunSlice() {
	//可变长数组
	var s1 []string
	fmt.Printf("s1: %v\n", s1)

	var s2 = make([]string, 3)
	s2 = []string{"a", "b", "c"}
	s2 = append(s2, "d")
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("len(s2): %v\n", len(s2))
	fmt.Printf("cap(s2): %v\n", cap(s2))

	//截取
	fmt.Printf("s2[1:2]: %v\n", s2[1:2])
	fmt.Printf("s2[3:]: %v\n", s2[3:])
}

//切片删除操作
func RunSliceDelete() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s2 = append(s1[:3], s1[6:]...)  //删除索引3-5元素
	var s3 = append(s1[:2], s1[4:6]...) //����
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
}
