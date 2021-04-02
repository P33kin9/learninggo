package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("completed, notification monitor stopped")
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "monitor exited, stopping...")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine is monitoring...")
			time.Sleep(2 * time.Second)
		}
	}
}
