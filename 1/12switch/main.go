package main

import "fmt"

// switch,简化大量判断
func main()  {
	a := 2
	if a == 1{
		fmt.Println("1")
	}else if a == 2{
		fmt.Println("2")
	}else if a == 3{
		fmt.Println("3")
	}else {
		fmt.Print("无效的输入")
	}

	// switch 简化上面代码
	switch b :=6; b {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3,4,5,6:
		fmt.Println("3")
	default:
		fmt.Println("无效的输入")
	}
}
