package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithSelectDemo() {
	fmt.Println("=== Context与Select结合示例 ===")

	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 42
	}()

	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-ctx.Done():
		fmt.Println("超时:", ctx.Err())
	}
	fmt.Println()
}

func main() {
	contextWithSelectDemo()
}
