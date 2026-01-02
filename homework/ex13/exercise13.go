package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	c1 := fetchPrice("shopee")
	c2 := fetchPrice("grab")
	c3 := fetchPrice("be")
	out := merge(c1, c2, c3)
	for v := range out {
		fmt.Println(v)
	}
}
func fetchPrice(source string) <-chan string {
	out := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ran := rand.IntN(3)
		out <- fmt.Sprintf("Giá %s: %d", source, ran)
		close(out)
	}()
	return out
}
func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	results := make(chan string)
	output := func(c <-chan string) {
		defer wg.Done()
		for n := range c {
			results <- n
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}
