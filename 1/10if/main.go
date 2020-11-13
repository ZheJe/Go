package main

import"fmt"

func ifDemo1() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}

	//fmt.Println(age)  // 在这里找不到age
}