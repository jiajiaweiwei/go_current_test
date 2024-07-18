package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("吃吃吃就知道吃")
}

//	func (a *Animal) Call() {
//		fmt.Println("Call-----")
//	}
func main() {
	a := Animal{}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
