package main

import (
	"fmt"
	"sync"
)

// MyObject 定义一个结构体作为对象池中存储的对象类型
type MyObject struct {
	Value int
}

func main() {
	// 创建一个对象池
	var pool = sync.Pool{
		New: func() interface{} {
			// 创建一个新的对象
			return &MyObject{}
		},
	}

	// 从对象池中获取对象
	obj := pool.Get().(*MyObject)
	fmt.Println("Object 1:", obj)

	// 设置对象的值
	obj.Value = 10
	fmt.Println("Object 1 after setting value:", obj)

	// 归还对象到对象池
	pool.Put(obj)
	obj2 := pool.Get().(*MyObject)
	fmt.Println("Object 2:", obj2)

	// 设置对象的值
	obj2.Value = 22
	fmt.Println("Object 2 after setting value:", obj2)

	// 归还对象到对象池
	pool.Put(obj)
	// 再次从对象池中获取对象
	obj3 := pool.Get().(*MyObject)
	fmt.Println("Object 3:", obj3)
	obj4 := pool.Get().(*MyObject)
	fmt.Println("Object 4:", obj4)
	// 注意：对象在归还到对象池后，原来的状态会被清除
}
