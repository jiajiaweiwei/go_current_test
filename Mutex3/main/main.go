package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
	"time"
)

func test1() {
	var lock1 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	count := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				lock1.Lock()
				count++
				lock1.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
func test2() {
	var mu sync.Mutex
	defer mu.Lock()
	fmt.Println("---test forget Unlock-----")
}

func main() {
	//var mu sync.Mutex
	//mu.Lock()
	//test3(mu)
	var rm RecursiveMutex

	// 示例1：递归调用
	fmt.Println("示例1：递归调用")
	rm.Lock()
	defer rm.Unlock()
	recursiveFunction(&rm, 3)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// 示例2：嵌套方法调用
	fmt.Println("示例2：嵌套方法调用")
	rm.Lock()
	defer rm.Unlock()
	nestedMethod(&rm)
}
func test3(mu sync.Mutex) {
	mu.Lock()
	fmt.Println("---test copy mutex---")
	mu.Unlock()
}

// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 记录当前持有锁的goroutine id
	recursion int32 // 记录goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get() // 使用goid库直接获取goroutine id
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		// 将重入次数+1
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get() // 使用goid库直接获取goroutine id
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		// 错误日志 输出 错误的 goroutine id 和正在占用锁的goroutine id
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
func recursiveFunction(rm *RecursiveMutex, x int) {
	if x <= 0 {
		return
	}
	rm.Lock()
	defer rm.Unlock()
	fmt.Printf("递归调用：%d\n", x)
	recursiveFunction(rm, x-1)
}

// 嵌套方法调用示例
func nestedMethod(rm *RecursiveMutex) {
	rm.Lock()
	defer rm.Unlock()
	fmt.Println("外部方法调用")
	internalMethod(rm)
}

func internalMethod(rm *RecursiveMutex) {
	rm.Lock()
	defer rm.Unlock()
	fmt.Println("内部方法调用")
}
