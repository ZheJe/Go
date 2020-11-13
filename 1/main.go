package main

import "fmt"

// 练习题
func main()  {
// 1,编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	a := 1
	b := 1.123123
	c := true
	d := "shi界"
	fmt.Printf("a:%T %v\n", a, a)
	fmt.Printf("b:%T %v\n", b, b)
	fmt.Printf("c:%T %v\n", c, c)
	fmt.Printf("d:%T %v\n", d, d)
// 编写代码统计出字符串"hello,世界"中汉字的数量。
	aa := "hello,世界"
	bb := []rune(aa)
	dd := 0
	
	for _, cc := range bb {
		if len(string(cc)) > 2 {
			dd++
		}
	}
	fmt.Print(dd)
}