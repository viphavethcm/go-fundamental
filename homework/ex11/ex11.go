package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	nums := make(chan int)
	results := make(chan int)
	errors := make(chan error)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Generator
	go generator(nums, ctx)

	// Start worker
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(nums, results, errors, &wg)
	}
	// Monitor Results closing in another goroutine
	go func() {
		wg.Wait()
		close(results)
	}()
	// Error handling
	go func() {
		for error := range errors {
			fmt.Println("🚨 Bắt được lỗi:", error)
			cancel()
		}
	}()
	result := accumulator(results)
	fmt.Println(result)
}
func generator(nums chan int, ctx context.Context) {
	for i := 1; i <= 100; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Generator: Bị hủy!")
			// Quan trọng: Phải đóng nums để các worker thoát vòng lặp range
			close(nums)
			return
		case nums <- i:
		}
	}
}

func worker(nums <-chan int, results chan int, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range nums {
		if num == 50 {
			select {
			case errors <- fmt.Errorf("Lỗi nghiêm trọng tại số 50"):
			default:
			}
			return
		}
		time.Sleep(10 * time.Millisecond)
		results <- num * num
	}
}

func accumulator(results <-chan int) int {
	result := 0
	for num := range results {
		result += num
	}
	return result
}
