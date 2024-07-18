package main

import (
	"fmt"
	"sync"
	"time"
)

// golang并发奠基之路专栏， 博客3
// mutex的底层实现 与历史发展
// mutex 使用的常见错误
func main() {
	fmt.Println("=========")
	//foo1()
	/*var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo2(c) // 复制锁*/
	// 测试重入的情况
	//l := &sync.Mutex{}
	//foo3(l)
	// 测试死锁
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2) // 需要派出所和物业都处理
	// 派出所处理goroutine
	go func() {
		defer wg.Done() // 派出所处理完成
		psCertificate.Lock()
		defer psCertificate.Unlock()
		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求物业的证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()
	// 物业处理goroutine
	go func() {
		defer wg.Done() // 物业处理完成
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()
		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求派出所的证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()
	wg.Wait()
	fmt.Println("成功完成")
	fmt.Println("=========")
}

// 测试常见使用错误
// 1. 加锁 和 解锁 没有成对出现
// 1.1 直接解锁一个为加锁的mutex会发生panic
func foo1() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world!")
}

// Counter 2.复制已经使用的mutex
type Counter struct {
	sync.Mutex
	Count int
}

// 这里Counter的参数是通过复制的方式传入的
func foo2(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

// 3.重入使用了mutex mutex不是可重入的锁，所以如果对mutex使用重入，会导致报错
func foo3(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}
func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

// 3.1 实现可重入锁的方案
// 3.1.1 使用协程id
// 3.1.2 给 协程添加token

// 4.死锁
// 死锁 也会导致mutex异常，
// 导致死锁的四大必要条件
// 1.互斥
// 2.持有和等待
// 3.不可剥夺
// 4.环路等待
