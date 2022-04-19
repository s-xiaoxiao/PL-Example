package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个协程都会运行该函数/
// 睡眠一秒钟，以此来模拟耗时的任务
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	fmt.Println("想要等待多个协程完成，我们可以使用 wait group")
	fmt.Println("wait groups")

	// 这个 WaitGroup 用于等待这里启用的所有协程完成。
	// 注意：如果WaitGroup 显式传递到函数中，则应使用 指针
	var wg sync.WaitGroup

	// 启动几个协程，并为其递增WaitGroup的计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 避免在每个协程闭包中重复利用相同i值
		i := i

		// 将worker调用包装在一个闭包中，可以确保通知WaitGroup此工作线程已完成。
		// 这样worker线程本身就不必知道其执行中涉及的并发原语。
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// 阻塞，直到waitGroup 计数器恢复为0；即所有协程的工作都已经完成。
	wg.Wait()
}
