package main

import (
	"fmt"
	"time"
)

func main() {
	// Go 的选择器（select）让你可以同时等待多个通道操作。
	// 		将协程、通道和选择器结合，是Go的一个强大特性。

	fmt.Println("select - 通道选择器")

	// 创建两个通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 各个通道将在一定时间后接收一个值，通过这种方式来模拟并行的协程执行时造成的阻塞
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 使用 select 关键字来同时等待这两个值，并打印各自接收到的值
	// 这个语法在例子里没有说明。
	// 应该是：select 会等待 协程结束时,通过通道写值进行匹配。然后进入下一个循环
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
