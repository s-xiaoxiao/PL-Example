package main

import (
	"fmt"
	"time"
)

// 这是 worker 程序,我们会在并发的运行多个 worker.
// worker 将在jobs 频道上接收工作,并在results上发送相应的结果.
// 每个worker 我们都会 sleep 一秒钟,以模拟一项昂贵的(耗时一秒钟的)任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("在这个例子中,我们将看到如何使用协程与通道实现一个工作池")
	fmt.Println("worker pools")
	fmt.Println()

	const numJobs = 5

	// 为了使用worker工作池并且收集其的结果,我们需要2个通道
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 这里启用了3个worker,初始是阻塞的,因为还没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送了5个jobs,然后close 这些通道,表示这些就是所有的任务了
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// 最后,我们收集所有这些任务的返回值.这也确保了所有的worker协程都已完成.
	// 另一个等待多个协程的方法是使用 WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// 运行程序,显示五个任务被多个worker执行.尽管所有的工作中共要花费5秒钟,
	// 但该程序只花了2秒钟,因为3个worker是并行的
}
