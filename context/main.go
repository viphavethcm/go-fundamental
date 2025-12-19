package main

import (
	"context"
	"fmt"
	"time"
)

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			name := ctx.Value("name")
			fmt.Println("Username: ", name)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "name", "duy")
	go employee(ctx)
	time.Sleep(3 * time.Second)
}
