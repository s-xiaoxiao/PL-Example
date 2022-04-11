package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Array, It's Arrays grammar")

	var a [5]int

	fmt.Println("emp:", a)

	a[4] = 100
	// 设置值
	fmt.Println("set:", a)
	// 获取值
	fmt.Println("get:", a[4])
	// 数组长度
	fmt.Println("len:", len(a))

	// 声明并且初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 二维数据
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
