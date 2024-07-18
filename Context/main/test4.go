package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有超时时间的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 启动一个并发任务，传递上下文
	go doWork1(ctx)
	// 等待一段时间，确保上下文超时
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine completed")
}
func doWork1(ctx context.Context) {
	// 模拟一些工作
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker goroutine received cancel signal")
			return
		default:
			// 模拟一些工作
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Working...")
		}
	}
}
