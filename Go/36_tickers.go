package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器是当你想要在未来某一刻执行一次时使用的
	// 打点器 则是为你想要以固定的时间间隔重复执行而准备的.
	fmt.Println("tickers")

	// 打点器和定时器的机制有点相似:使用一个通道来发送数据.
	// 这里我们使用通道内存内建的select,等待每500ms到达一次的值
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 我们将在1600ms 后停止这个打点器
	time.Sleep(1600 * time.Millisecond)

	// 打点器和定时器一样被停止.打点器一旦停止,将不能再从它的通道中接收到值.
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
