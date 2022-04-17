package main

import "fmt"

func main() {
	// 非阻塞通道操作

	// 常规的通过通道发送和接收数组是阻塞的,
	// 然而,我们可以使用带一个的default 子句的select来实现非阻塞的发送 接收.
	// 甚至是非阻塞的多路 selects

	messages := make(chan string)
	signals := make(chan bool)

	// 这是一个非阻塞接收的例子.如果messages中存在,然后select将这个值带入<-messages case中
	// 否则,就直接到default分子中
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞发送的例子,代码结构和上面接收的类似.
	// msg不能被发送到messages通道中,因为这是个无缓冲区通道,
	// 并且也没有接收者,因此 default 会执行
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在default 前使用多个case子句来实现一个多路的非阻塞的选择器
	// 这里我们试图在message和signals上 同时使用非阻塞的接收操作
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
