package main

import (
	"fmt"
)

func main() {
	// for 是Go 中唯一的循环结构。这里会展示for 循环的一些基本使用方式。

	// most base method
	// 最基础的方式,单个循环条件
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 初始 条件 后续
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 重复执行; 遇到break and return end
	for {
		fmt.Println("loop")
		break
	}

	// continue

	for n := 0; n <= 5; n++ {

		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

	// 以及 range channels数据结构
}
