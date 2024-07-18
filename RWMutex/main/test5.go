package main

import (
	"fmt"
	"time"
)

var (
	num  = make(chan struct{})
	word = make(chan struct{})
)

func main() {

	go numP()
	go wordP()
	num <- struct{}{}
	time.Sleep(2 * time.Second)
}
func numP() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Println(1)
		word <- struct{}{}
	}
}
func wordP() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Println("A")
		num <- struct{}{}
	}
}
