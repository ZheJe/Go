package main

import "fmt"

// Go语言中推荐使用驼峰式明明
// var student_name string
var studentName string
// var studentName string

// 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)
// s5 := "hehe"

func main() {
	name = "理想"
	age = 16
	isOk = true
	studentName = "test"
	// Go语言变量声明必须使用,不使用就编译不过去
	fmt.Print(isOk)             // 在终端中输出要打印的内容
	fmt.Printf("name:%s", name) // %s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)			// 打印完指定的内容后悔加一行换行符
	fmt.Println(studentName)
	// heiheihei = "嘿嘿嘿"

	// 声明变量同事赋值
	var s1 string = "web"
	fmt.Println(s1)
	// 类型推导 (根据值判断该变量是什么类型,推荐这样)
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明 (必须在函数内部使用)
	s3 := "哈哈哈"
	fmt.Println(s3)
	// s1 := "10"  //同一个作用于({}) 中不能重复声明相同的变量
	// 匿名变量是一个特殊的变量:_(后面学函数会学)
}
