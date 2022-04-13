package main

import (
	"fmt"
)

// n 不等于一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("recursion 递归")

	fmt.Println(fact(7))

	// 闭包也可以是递归的。
	// Note 没有明白闭包也可以是递归的这句话的意思 2022.4.14(周四)。
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
