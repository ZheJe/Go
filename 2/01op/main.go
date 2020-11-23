/*
 * @名字: 喆劼
 * @Date: 2020-11-16 21:28:25
 * @LastEditTime: 2020-11-17 09:19:52
 * @Description: 
 */
package main

import "fmt"

func main()  {

	// 1, 求数组[1, 3, 5, 7, 8]所有元素的和
	var sum int
	a := [...]int{1, 3, 5, 7, 8}
	for i, _ := range a {
		sum += a[i]
	}
	fmt.Println(sum)
	// 2, 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	a2 := [...]int{1, 3, 5, 7, 8}
	for i, _ := range a2 {
		if i == 0 || i == 1 {
			for i2, _ := range a2{
				if i2 == 3 || i2 == 2{
					if a[i] + a[i2] == 8{
						fmt.Println(i,a2[i],"|",i2,a2[i2])
					}
				}
			}
		}
	}
}