package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var key string

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}
