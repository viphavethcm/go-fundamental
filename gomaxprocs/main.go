package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("So CPU: %d\n", numCPU)

	runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go heavyTask(&wg)
	}
	wg.Wait()
	fmt.Printf("Tong thoi gian: %s\n", time.Since(start))
}

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 0; i < 100e8; i++ {
		sum += 1
	}
}
