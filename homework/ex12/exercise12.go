package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	queue := make(chan string)
	var wg sync.WaitGroup
	go generator(queue)
	wg.Add(1)
	go worker(queue, &wg)
	timeoutSignal := time.After(3 * time.Second)

	// Main đứng đây chờ tín hiệu từ channel này
	<-timeoutSignal
	fmt.Println("--- ⚠️  ADMIN: Nhận tín hiệu tắt (từ time.After)... ---")

	// Các bước sau giữ nguyên
	wg.Wait()
	fmt.Println("--- ✅ Server Stopped Gracefully ---")
}

func generator(queue chan<- string) {
	for i := 1; i <= 20; i++ {
		queue <- "email " + strconv.Itoa(i) + "."
	}
	close(queue)
}

func worker(queue <-chan string, wg *sync.WaitGroup) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	for message := range queue {
		<-ticker.C
		fmt.Println("Worker: Sending message:", message)
	}
}
