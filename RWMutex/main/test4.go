package main

import (
	"fmt"
	"time"
)

var (
	cat  = make(chan struct{})
	fish = make(chan struct{})
	dog  = make(chan struct{})
)

func main() {
	for i := 0; i < 100; i++ {
		go printfCat()
		go printfFish()
		go printfDog()
	}
	cat <- struct{}{}
	time.Sleep(2 * time.Second)
}

// 方法2 使用通道
func printfCat() {
	<-cat // 从 cat 通道接收信号，此处会阻塞直到有数据可接收
	fmt.Println("cat")
	fish <- struct{}{}
}
func printfFish() {
	<-fish
	fmt.Println("fish")
	dog <- struct{}{}
}
func printfDog() {
	<-dog
	fmt.Println("dog")
	cat <- struct{}{}
}
