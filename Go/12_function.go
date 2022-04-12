package main

import "fmt"

// 明确标识接受两个 int 参数，返回值为int 返回值需要明确
func plus(a int, b int) int {
	return a + b
}

// 当参数连续相同时，只需声明最后一个参数即可，可以忽略之前的
func plusPlus(a, b, c int) int {
	return a + b + c

}
func main() {
	// 函数是 Go的核心

	// 通过函数名(参数列表)来调用函数
	res := plus(1, 2)

	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)

	fmt.Println("1+2+3=", res)

}
