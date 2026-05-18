package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	go slowOperation(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("MAIN DONE")
}
func slowOperation(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("OPERATION COMPLETED")

	case <-ctx.Done():
		fmt.Println("Operation canceled", ctx.Err())
	}
}
