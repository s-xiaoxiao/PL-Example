package main

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num= %b", b.num)
}

type container struct {
	base
	str string
}

func main() {
	fmt.Println("embedding")
	// Go 支持对于结构体（struct)和接口（interfaces)的嵌入（embedding）
	// 以表达一种更加无缝的组合（composition）类型

	// container 嵌入了一个base,当创建含有嵌入的结构体，必须对嵌入进行显式的初始化
	// 在这里使用嵌入的类型当作字段的名字
	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	// 我们可以直接在 co 上访问 base 定义的字段 例如：co.num
	fmt.Printf("co = {num %v,str:%v}\n", co.num, co.str)

	// 或者使用完整路径
	fmt.Println("also num:", co.base.num)
	// 由于container嵌入了base，因此base的方法也成为了container的方法。可以直接调用
	fmt.Println("describe:", co.describe())

	// 可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。
	// 因为嵌入了base 在这里我们看到container也实现了 describer接口
	type describer interface {
		describe() string
	}

	var d describer = co

	fmt.Println("describer", d.describe())
}
