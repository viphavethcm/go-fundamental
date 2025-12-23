package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	slowWorker(ctx)
}
func slowWorker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Worker: Sếp bắt dừng! Lý do:", ctx.Err())
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Worker: Xong việc (nhưng chắc chắn không chạy tới đây được)")
	}
}
