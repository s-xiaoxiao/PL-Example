package main

import "fmt"

func main() {
	// 关闭一个通道意味着不能再向这个通道发送值了.
	// 该特性可以向通道的接收方传达工作已经完成的信息.

	// example:
	// 在这个例子中,我们将使用一个jobs通道,将工作内容,从 main()协程传递达到一个工作协程中.
	// 当我们没有更多的任务传递给工作协程时,我们将 close 这个jobs通道

	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作协程.使用 j,more:= <- jobs 循环的从jobs接收数据.
	// 根据第二个值都已经接收完毕,那么more值将是false.
	// 当我们完成所有任务时,会使用这个特性通过done通道通知main协程
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 使用jobs发送三个任务到工作协程中,然后关闭jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 通道同步返回等待任务结束
	<-done
}
