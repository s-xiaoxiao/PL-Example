package main

import (
	"fmt"
)

// ping 函数定义了一个只能发送（只写）数据的通道。尝试从这个通道接收数据会是一个编译错误
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接收两个通道，
// ping仅用于接收数据（只读），
// pong用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	fmt.Println("channel directions")
	fmt.Println("通道方向")
	fmt.Println("------------------")
	// 当使用通道作为函数的参数时，你可以指定这个通道时候为只读或者只写。
	// 该特性可以提升程序的类型安全

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
