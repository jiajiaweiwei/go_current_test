package main

import (
	"context"
	"fmt"
	"time"
)

//使用 WithCancel 和 WithValue 写一个级联的使用 Context 的例子，验证一下 parent
//Context 被 cancel 后，子 conext 是否也立刻被 cancel 了。
// 测试

func main() {
	// 创建一个根上下文
	rootCtx := context.Background()
	// 使用 WithCancel 创建一个取消上下文和取消函数
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	// 在根上下文中添加一个值
	ctxWithValue := context.WithValue(ctx, "key1", "value1")
	// 启动一个 goroutine 来处理任务
	go func(ctx context.Context) {
		// 在子上下文中进行工作
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Child context received cancel signal")
				return
			default:
				// 获取并打印上下文中的值
				value := ctx.Value("key1")
				fmt.Printf("Working with value: %v\n", value)

				// 模拟一些工作
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctxWithValue)
	// 模拟一段时间后取消根上下文
	time.Sleep(1 * time.Second)
	cancel()
	// 等待 goroutine 执行完成
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main goroutine completed")
}
