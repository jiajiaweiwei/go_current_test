package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个长时间任务
func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		// 模拟任务需要5秒钟
		fmt.Println("Task completed")
	case <-ctx.Done():
		// 如果 context 被取消，打印取消原因
		fmt.Println("Task cancelled:", ctx.Err())
	}
}

func main() {
	// 创建一个有截止时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 确保在 main 函数退出时取消 context

	// 检查 context 的截止时间
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("Deadline set for:", deadline)
	} else {
		fmt.Println("No deadline set")
	}

	// 启动长时间任务
	go longRunningTask(ctx)

	// 等待任务完成或 context 被取消
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Main function completed")
	case <-ctx.Done():
		fmt.Println("Main function cancelled:", ctx.Err())
	}
}
