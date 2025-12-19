package main

import (
	"fmt"
	"sync"
)

func main() {
	token := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			token++
			wg.Done()
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(token)
}
