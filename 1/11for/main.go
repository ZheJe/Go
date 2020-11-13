package main

import "fmt"

func main()  {
	for i := 1; i <= 9; i++ {
		for z := 1; z <= i; z++ {
			fmt.Printf("%v*%v=%v\t", i, z, i*z)
		}
		fmt.Println()
	}
	for i := 9; i >= 1; i-- {
		var flag = false
		for z := 1; z <= i; z++ {
			if i == 8 {	// 当i=8跳过,继续下一次循环
				continue
			}
			if i == 5 {	//当i=5退出循环
				flag = true
				break
			}
			fmt.Printf("%v*%v=%v\t", z, i, i*z)
		}
		if flag {
			break // 跳出for循环(外层的for循环)
		}
		fmt.Println()
	}
}
