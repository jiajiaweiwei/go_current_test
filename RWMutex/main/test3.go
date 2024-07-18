package main

import (
	"fmt"
	"sync"
)

var (
	mu       sync.Mutex
	turn     = 0
	maxTurns = 300 // 每个函数打印次数
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go printCat(&wg)
	go printFish(&wg)
	go printDog(&wg)

	wg.Wait()
	fmt.Println("over-----")
}

func printCat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < maxTurns; i++ {
		mu.Lock()
		for turn%3 != 0 {
			mu.Unlock()
			mu.Lock()
		}
		fmt.Println("cat")
		turn++
		mu.Unlock()
	}
}

func printFish(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < maxTurns; i++ {
		mu.Lock()
		for turn%3 != 1 {
			mu.Unlock()
			mu.Lock()
		}
		fmt.Println("fish")
		turn++
		mu.Unlock()
	}
}

func printDog(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < maxTurns; i++ {
		mu.Lock()
		for turn%3 != 2 {
			mu.Unlock()
			mu.Lock()
		}
		fmt.Println("dog")
		turn++
		mu.Unlock()
	}
}
