package main

import "fmt"

// byte和rune类型

// Go语言中为了处理费ASCII码类型的字符 定义了新的run类型

func main()  {
	s := "hello 世界 세계"
	// len()求的是byte字节的数量
	n := len(s)  // 求字符串s的长度,把长度保存到变量n中
	fmt.Println(n)

	// for i := 0 ; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// 	fmt.Printf("%c\n", s[i])  // %c:字符
	// }
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	// "Hello" => 'H' 'e' 'l' 'l' 'o'
	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
}