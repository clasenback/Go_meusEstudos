package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)

	fmt.Println("----------")

	ctxTodo := context.TODO()
	fmt.Println("context:\t", ctxTodo)
	fmt.Println("context err:\t", ctxTodo.Err())
	fmt.Printf("context type:\t%T\n", ctxTodo)

	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)

	fmt.Println("----------")

	cancel()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)

	fmt.Println("----------")

	t := time.Duration(rand.Intn(2000)) * time.Millisecond
	fmt.Println("Tempo de processamento:", t)
	ctx, cancel = context.WithTimeout(context.Background(), t)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("context err:\t", ctx.Err()) // prints "context deadline exceeded"
	}

	fmt.Println("----------")

	tt := rand.Intn(2000)
	d := time.Now().Add(time.Duration(tt) * time.Millisecond)
	fmt.Println("Deadline:", d, time.Duration(tt*1000000))
	ctx, cancel = context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	fmt.Println("----------")
}
