package main

import "fmt"

// map

func main()  {
	var m1 map[string]int
	fmt.Println(m1 == nil) // 还没有初始化(在内存开辟空间)
	m1 = make(map[string]int, 10) // 要估算好该map容量,避免在程序运行间再次扩容
	m1["理想"] = 18
	m1["现实"] = 1
	fmt.Println(m1)
	// 约定成俗用ok接受返回的布尔值
	fmt.Println(m1["天基"]) // 如果不存在这个key就会返回0
	value, ok := m1["天基"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}