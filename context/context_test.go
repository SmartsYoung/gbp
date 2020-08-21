package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_watch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)

}

func Test_work(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	defer cancel()
	t.Log("Hey, I'm going to do some work\n")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	t.Log("Finished. I'm going home\n")

}

func Test_WithDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
