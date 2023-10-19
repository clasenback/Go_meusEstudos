package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		var t time.Duration
		for {
			t = time.Millisecond * time.Duration(rand.Intn(200))
			time.Sleep(t)
			select {
			case <-ctx.Done():
				return
			default:
				n++
				fmt.Println("working", n, "in", t)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}
