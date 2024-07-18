package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有取消功能的上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	// 启动一个并发任务，传递上下文
	go doWork(ctx)
	// 模拟一段时间后取消操作
	time.Sleep(1 * time.Second)
	cancel() // 取消操作
	// 等待任务完成
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine completed")
}
func doWork(ctx context.Context) {
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
