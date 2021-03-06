/*
 * @名字: 喆劼
 * @Date: 2020-11-16 21:06:11
 * @LastEditTime: 2020-11-16 21:28:57
 * @Description: 
 */
package main

import "fmt"

func main()  {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T", a1, a2)
	
	// 数组初始化
	// 如果不出适合,默认元素都是零值 (布尔值: flase, 整型与浮点型都是0,字符串:"")
	// 1, 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 初始化方式2
	// a10 := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	// 3,初始化方式3,根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1, 根据索引遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1, 2] [3, 4] [5, 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1,2,3}		// [1 2 3]
	b2 := b1				// [1 2 3] 复制粘贴
	b2[0] = 100  			// b2:[100 2 3]
	fmt.Println(b1, b2)		//b1:?
}