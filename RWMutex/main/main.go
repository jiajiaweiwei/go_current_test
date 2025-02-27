package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var counter Counter
	//for i := 0; i < 10; i++ { // 10个reader
	//	go func() {
	//		for {
	//			fmt.Println("--读取到的数据：--", counter.Count(), "--读取到的数据：--") // 计数器读操作
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//for { // 一个writer
	//	counter.Incr() // 计数器写操作
	//	time.Sleep(time.Millisecond * 10)
	//}
	var mu sync.RWMutex
	// writer,稍微等待，然后制造一个调用Lock的场景
	go func() {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}()
	go func() {
		factorial(&mu, 10) // 计算10的阶乘, 10!
	}()

	select {}
}

// 递归调用计算阶乘
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 { // 阶乘退出条件
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n // 递归调用
}
