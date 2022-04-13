package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// Go 支持 指针，允许在程序中通过 引用传递 来传递值和数据结构。

	// zeroval 是值传递 ival 会得到实参的一个拷贝
	// zeroptr 是指针传递 ival 将被赋值为 0

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 &i语法来取得 i的内存地址。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也可以被打印
	fmt.Println("pointer:", &i)
}
