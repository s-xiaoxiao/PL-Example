package main

import (
	"fmt"
	"time"
)

func main() {
	// 我们经常需要在未来的某个时间点运行Go代码,或者每隔一定时间重复运行代码.
	// Go 内置的定时器和打点器特性让这些变得很简单.
	// 我们会先学习定时器,然后再学习打点器

	// 定时器标识在未来某一时刻的独立事件.
	// 你告诉定时器需要等待的时间,然后它将提供一个用于通知的通道.这里的定时器将等待2秒
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer.c 会一直阻塞,直到定时器通知C明确发送了定时器失效的值.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 如果你需要的仅仅是单纯的等待,使用time.Sleep就可以了.
	// 使用定时器的原因之一就是,你可以在定时器触发之前将其取消
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 给Timer2足够时间来触发它,以证明它实际上已经停止了
	time.Sleep(2 * time.Second)

}
