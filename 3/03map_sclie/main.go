package main

import "fmt"

// map与sclie结合

func main()  {
	// 元素类型的切片
	var s1 = make([]map[int]string, 10, 10)
	// 没有对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "A"
	fmt.Println(s1)

	//值为切片的map
	var m1 = make(map[string][]int, 10)
	m1["天津"] = []int{10, 20, 30}
	fmt.Println(m1)
}