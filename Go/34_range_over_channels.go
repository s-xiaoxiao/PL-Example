package main

import "fmt"

func main() {
	// Go 提供了for和range为基本的数据结构提供了迭代的功能。
	// 我们也可以使用这个语法来遍历的从通道中取值

	fmt.Println("range over channels")
	fmt.Println("通道遍历")
	fmt.Println()

	queue := make(chan string, 2)

	// 我们将遍历在 queue 通道中的两个值
	queue <- "one"
	queue <- "two"
	close(queue)

	// range 迭代从queue中得到每个值。因为我们在前面close了这个通道。
	// 所以，这个迭代会在接受完2个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}

	// 这个例子也让我们看到，一个非空的通道也是可以关闭的，
	// 并且,通道中剩下的值仍然可以被接收到.
}
