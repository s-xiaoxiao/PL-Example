package main

import (
	"fmt"
)

// 这个函数可以接受任意数量的int作为参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {

	// 可变参数函数 ：例如 fmt.Println

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 可以将数组转换为可变参数
	sum(nums...)
}
