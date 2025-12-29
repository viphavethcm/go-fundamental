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
	go generator(nums, ctx)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(nums, results, errors, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	go func() {
		for error := range errors {
			errors <- fmt.Errorf("error: %s", error.Error())
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
			return
		case nums <- i:
		}
	}
	close(nums)
}

func worker(nums <-chan int, results chan int, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range nums {
		if num == 50 {
			errors <- fmt.Errorf("Lỗi nghiêm trọng tại số 50")
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
