package main

import "fmt"
import "strings"
// 转译

func main()  {
	path := "\"C:\\Users\\dusho\\Documents\\Go\\Go\\1\\08string\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行的字符串
	s2 := `
世情薄
				人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `C:\Users\dusho\Documents\Go\Go\1\08string`
	fmt.Println(s3)

	// 字符串相关操作
	name := "理想"
	world := "dd"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	// fmt.Printf("%s%s", name, world)  //  与上面区别,Sprint可以保存到变量,这个是直接打印出来
	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcde"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}