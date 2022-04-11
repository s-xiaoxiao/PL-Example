package main

import (
	"fmt"
)

func main() {

	// map 是 Go内建的关联数据类型，它提供了一种存储键值对的方法。
	// ( 其他语言中也被称为 哈希(hash) 或 字典(dict) )

	//  make(map[key-type]value-type)

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// return key/value 数量
	fmt.Println("len:", len(m))

	// 删除 键值对
	delete(m, "k2")

	// map 取值时,可以选择是否接受第二个返回值,表明map中是否存在该键。
	// _空白标识符，对于不需要该值的语句，可以使用 _ 来代替。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 创建 map 并且初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
