package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// NOTE: It is handy to create context at any level.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 3*time.Second, "hello!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
