package main

import "fmt"

// 函数是一段代码的风闸UN个
// 把一段逻辑抽象出来封装到一个函数中,给他起个名字,每次用到它直接用函数名调用就可以
// 使用函数能够让带吗结构更清晰,更简洁

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值

func f1(x int,  y int)  {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2()  {
	fmt.Println("f2")
}

//没有参数担忧返回值
func f3()  int {
	return 3
}

// 返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return  // 使用命名返回值可以省略return后面
}

// 多个返回值
func f5() (int, string)  {
	return 1, "西青"
}

// 参数的类型简写:当参数中连续多个参数类型一致时,我们可以将非最后IG额参数的类型省略
func f6(x, y, z int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数, 必须方赞函数参数的最后
func f7(x string, y ...int)  {
	fmt.Println(x)
	fmt.Println(y) // y的类型切片[]int
}

// Go语言中函数没有默认参数这个概念

func main()  {
	r := sum(1, 2)	
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f7("下雨了")
	f7("下雨了", 1, 2, 3, 4, 5, 6)
}