package main

import "fmt"

// 运算符

func main()  {
	var (
		a = 5
		b = 2
	)
	// 算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++  // 单独的语句,不能放在=的右边赋值  ==> a = a + 1
	b++	 // 单独的语句,不能放在=的右边赋值  ==> b = b + 1
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b)  // Go预言师强类型,相同的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a > b)

	//逻辑运算符
	//如果年龄大于18岁并且小于60
	if age := 22; age > 18 && age < 60 {
		fmt.Println("苦逼上班的!")
	} else {
		fmt.Println("不用上班!")
	}
	// 如果年龄小于18岁或者年龄大于60岁
	if age := 22; age > 18 || age < 60 {
		fmt.Println("苦逼上班的!")
	} else {
		fmt.Println("不用上班!")
	}

	// not取反.原来为真就为假,原来为假就位真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	// 位运算:针对的是二进制数
	// 5的二进制: 101
	// 2的二进制: 10
	// &:按位与 (两位均为1才为1)
	fmt.Println(5 & 2)  // 000
	// |:按位活 (两位有1个位1就为1)
	fmt.Println(5 | 2)  // 111  :  7
	// ^:按位异或 (两位不一样则为1)
	fmt.Println(5 ^ 2)  // 111  :  7
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010  =>  10
	fmt.Println(1 >> 10)  // 10000000000  >=  1024
	// >>:将二进制位有移指定位数
	fmt.Println(5 >> 1)  // 10  =>  12
	var m = int8(1)  	 // 只能存8位
	fmt.Println(m << 10) // 10000000000

	// 数值运算符,用来给变量赋值的
	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2
	
	x <<= 2 // x = x << 2
	x &= 2 // x = x & 1
	x |= 2 // x = x | 1
	fmt.Println(x)

}