package main

import "fmt"

// intSeq 函数返回一个在函数体内定义的匿名函数。
// 返回的函数使用闭包的方式隐藏变量i。返回的函数隐藏变量i以形成闭包
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	// closures 含义为闭包
	// Go 支持 匿名函数，并能用其构造闭包。想定义一个不需要命名的内联函数时是很实用的

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()

	fmt.Println(newInts())
}
