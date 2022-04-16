package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 携程（goroutine） 是轻量级的执行线程。
	fmt.Println("goroutine")

	f("direct")

	// 使用 go 调用这个函数，这个新的Go 协程将会并发地执行这个函数
	go f("goroutine")

	// 匿名函数启用一个协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在两个协程在独立的协程中 异步地运行，然后等待两个协程完成（更好的方法是使用 WaitGroup）
	time.Sleep(time.Second)
	fmt.Println("done")
}
