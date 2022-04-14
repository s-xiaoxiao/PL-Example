package main

import (
	"fmt"
)

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))

	for k := range m {
		r = append(r, k)

	}
	return r
}
// 作为泛型类型的示例,List是一个具有任意类型值得单链表
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 我们可以像在常规类型上一样定义泛型类型的方法
// 但我们必须保留类型参数.这个类型式 List[T]
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
func main() {

	// 泛型
	// 从1.18版本开始，Go添加了对泛型的支持，也即类型参数。

	var m = map[int]string{1: "2", 2: "4", 4: "8"}


	// 调用泛型函数的时候,我们经常可以使用类型推断.
	// 注意当调用Mapkeys的时候.我们不需要为K和V指定类型-编译器会进行自动推断
	fmt.Println("keys m:", MapKeys(m))

	//虽然也可以明确指定这些类型. 
	_ = MapKeys[int, string](m)

	
	// 单链表泛型例子
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.GetAll())
}
