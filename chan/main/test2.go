package main

import (
	"fmt"
	"net/http"
	"time"
)

type Student struct {
	id   string
	name string
}

type StudentReq struct {
	id   string
	name string
}

func main() {
	test()
}

func test() {

	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		client := http.Client{}
		request, _ := http.NewRequest("GET", "https://juejin.cn/post/7391699542007169033", nil)
		request.Header.Set("Content-Type", "application/json")
		response, _ := client.Do(request)
		defer response.Body.Close()
		fmt.Println(i, "-----:", +response.StatusCode)
	}
}
