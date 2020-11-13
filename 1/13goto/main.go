package main

import "fmt"

func main()  {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break  // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break  // 跳出for循环(外层的for循环)
		}
	}

	// Go+label实现跳出多层for 循环
	for i := 0; i < 10; i++      {
		for j := 'A'; j < 'Z'; j++{
			if j == 'C' {
				goto XX //跳到指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	XX : // lebel标签
		fmt.Println("over")
}
}