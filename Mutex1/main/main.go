package main

import (
	"fmt"
	"sync"
)

// golang并发奠基之路专栏， 博客2
func main() {
	CounterStructByMutest()
	CounterStructByMutest()
	CounterStructByMutest()
}

func testNoMutex() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func testMutex() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0
	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)
	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Lock() {
	//写具体的业务需求
	c.mu.Lock()
}
func (c *Counter) UnLock() {
	//写具体的业务需求
	c.mu.Unlock()
}
func CounterStructByMutest() {
	// 计数器的值
	var count Counter
	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)
	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				count.Lock()
				count.count++
				count.UnLock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.count)
}
