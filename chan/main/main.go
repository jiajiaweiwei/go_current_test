package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func() {
		fmt.Println("---开始接收---")
		for {
			// 从通道 c 中接收值，并赋给 temp
			temp, ok := <-c
			if !ok {
				fmt.Println("通道已关闭")
				return
			}
			fmt.Println(time.Now(), "接收到", temp)
		}
	}()
	// 等待一段时间，以便 goroutine 开始接收
	time.Sleep(1 * time.Second)
	// 向通道 c 中发送三个整数
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
	// 等待一段时间，以便 goroutine 完成接收
	time.Sleep(10 * time.Millisecond)
}
