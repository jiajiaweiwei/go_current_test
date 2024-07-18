package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func test(m *Mutex) {
	m.Lock()
	fmt.Println(atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))))
}
func test2(m *Mutex) {
	//m.Lock()
	fmt.Println(atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))))
}
func test3(m *sync.Mutex) {
	m.Lock()
	fmt.Println(atomic.LoadInt32((*int32)(unsafe.Pointer(&m))))
}
func test4(m *sync.Mutex) {
	//m.Lock()
	fmt.Println(atomic.LoadInt32((*int32)(unsafe.Pointer(&m))))
}
func testA() {
	var m Mutex
	test(&m)
	m.Unlock()
	test2(&m)
}
func testB() {
	var m sync.Mutex
	test3(&m)
	m.Unlock()
	test4(&m)
}
func main() {
	//testA()
	//fmt.Println("-----------")
	//testB()
	//fmt.Println("-----------")
	//try()
	fmt.Println("-----------")
	count()

}

// 实现一个Mutex定义的变量
const ( // 判断mutex状态的掩码常量
	// 将mutex源码中的掩码常量复制过来使用
	mutexLocked      = 1 << iota // 加锁标识位置  左移标识枚举值翻倍递增
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的其实bit位置
)

// Mutex 扩展一个Mutex结构
type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁
	// 采用直接将state置1的操作抢占锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	// 如果处于唤醒、加锁或者饥饿状态 这次请求就不参与竞争了，返回false
	// 原子方式读取state值
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}
	// 尝试在竞争状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
func try() {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	ok := mu.TryLock() // 尝试获取到锁
	if ok {            // 获取成功
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}
	// 没有获取到
	fmt.Println("can't get the lock")
}

// Count
// 获取state字段的值
func (m *Mutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift //得到等待者的数值
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}

// IsLocked 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// IsWoken 是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// IsStarving 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
func count() {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)

	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

// SliceQueue 自定义并发安全的队列
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
