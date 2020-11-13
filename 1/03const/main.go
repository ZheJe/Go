package main

import "fmt"


// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明,如果某一行声明没有复制,默认与上一行一样
const (
	n1 = 100
	n2
	n3 = 200
	n4
)

// iota:类似枚举
const (
	m1 = iota // 0
	m2		  // 1
	m3		  // 2
)

// 2会赋值给匿名
const (
	b1 = iota // 0
	b2 = iota // 1
	_  = iota // 2
	b3 = iota // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c4		  // 100
	c3 = iota // 3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)  
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//pi = 123
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("n4:", n4)

	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	fmt.Println("m3:", m3)
	
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}