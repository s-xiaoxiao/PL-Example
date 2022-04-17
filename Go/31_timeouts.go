package main

import (
	"fmt"
	"time"
)

func main() {

	// 超时对于一个需要连接外部资源，或者有耗时较长的操作的程序而言是很重要的。
	// 得益于通道和select,在Go中实现超时操作是简洁而优雅的

	c1 := make(chan string, 1)

	// 在这个例子中，假如我们执行一个外部调用，并在2秒后使用通道c1返回它的执行结果
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 使用select 实现一个超时操作。
	// res:=<-c1等待结果，<-time.After等待超时(1s)以后发送的值。
	// 由于select 默认处理第一个已准备好的接收操作,因此如果操作耗时超过了允许的1秒的话,将会执行超时case
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 如果我们允许一个长一点的超时时间:3s.就可以成功的从c2接收到值,并且打印出结果
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
