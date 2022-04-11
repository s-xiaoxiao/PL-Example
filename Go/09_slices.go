package main

import (
	"fmt"
)

func main() {

	// 创建一个slice 类型string 长度为3
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// slice 可以使用append 来增加元素;append 可以接受任意多个参数;并且返回一个新slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice can copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片操作： [low:high] 取出low到high的元素;不包含high
	// a b c d e f
	l := s[2:5]
	fmt.Println("sl1", l) // c d e

	l = s[:5]
	fmt.Println("sl2", l) // a b c d e

	l = s[2:]             // high 默认 len(s)
	fmt.Println("sl3", l) // c d e f

	// 声明并初始化一个 slice 变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slice 多维数据结构
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
