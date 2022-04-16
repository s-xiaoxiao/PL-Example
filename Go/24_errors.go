package main

import (
	"errors"
	"fmt"
)

// 按照惯例，错误通常是最后一个返回值并且是error类型。它是一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 给定错误信息一个基本的error值
		return -1, errors.New("can't work with 42")
	}

	return arg + 2, nil // 返回 nil 代表没有错误
}

type argError struct {
	arg  int
	prob string
}

// 可以通过实现 Error 方法来自定义error 类型。
// 这里使用自定义错误类型来表示上面例子中的参数错误。
func (e *argError) Error() string {
	return fmt.Sprintf("%d --- %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 在这个例子中，我们使用&argError语法来建立一个新的结构体，并提供了arg和prob两个字段的值
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	//Go

	// 测试 f1
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// 测试 f2
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// 自定义错误类型的数据，你需要通过类型断言来得到这个自定义错误类型的实例
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
