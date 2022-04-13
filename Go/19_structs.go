package main

import (
	"fmt"
)

// person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

// newPerson 使用给定的名字构造一个新的person 结构体
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p // 返回指定局部变量的指针。因为局部变量将在函数的作用域中继续存在
}

func main() {
	// Go 的结构体是带类型的字段集合。

	fmt.Println("结构体")

	// 使用这个语法创建新的结构体元素
	fmt.Println(person{"Bob", 20})

	// 初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})

	// 省略的被初始化为零值
	fmt.Println(person{name: "Fred"})

	// &前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 在构造函数中封装创新新的结构实例是一种习惯用法
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// 使用 . 来访问结构体字段
	fmt.Println(s.name)

	sp := &s
	// 也可以对结构体指针使用. 指针会被自动解引用
	fmt.Println(sp.age)

	// 结构体是可变的
	sp.age = 51
	fmt.Println(sp.age)
}
