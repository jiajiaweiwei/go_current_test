package main

import "fmt"

// 只能发送的通道
func sender(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i // 发送数据到通道
	}
	close(c) // 关闭通道，表示发送结束
}

// 只能接收的通道
func receiver(c <-chan int) {
	for num := range c {
		fmt.Println("Received:", num) // 从通道接收数据
	}
}

func main() {
	// 创建一个只能发送的通道
	sendOnly := make(chan<- int)
	// 创建一个只能接收的通道
	receiveOnly := make(<-chan int)

	// 在 Goroutine 中发送数据到 sendOnly 通道
	go sender(sendOnly)
	// 在 Goroutine 中接收数据从 receiveOnly 通道
	go receiver(receiveOnly)

	// 等待通道的处理
	fmt.Println("Main goroutine waiting...")
	// 这里可以添加更多的逻辑或者使用 select 等待 Goroutine 完成

	// 通道的使用结束后，记得关闭通道
	close(sendOnly)
	// 此处不关闭 receiveOnly，因为它是由 receiver 函数负责关闭的

	// 等待所有 Goroutine 结束
	// ...

	fmt.Println("Main goroutine completed")
}
