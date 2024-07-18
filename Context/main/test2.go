package main

import (
	"context"
	"fmt"
)

type contextKey1 struct{}
type contextKey2 struct{}

var UserKey1 = &contextKey1{}
var UserKey2 = &contextKey2{}

func main() {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, UserKey1, "UserKey1")
	ctx = context.WithValue(ctx, UserKey2, "UserKey2")
	fmt.Println(ctx.Value(UserKey1))
}
