package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var inventory = &Inventory{ticket: 100}

func main() {
	users := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, users, &wg)
	}
	go func() {
		for i := 1; i <= 500; i++ {
			users <- i
		}
		close(users)
	}()
	wg.Wait()
}

type Inventory struct {
	ticket int
	mu     sync.Mutex
}

func (inventory *Inventory) BuyTicket() bool {
	inventory.mu.Lock()
	defer inventory.mu.Unlock()
	time.Sleep(5 * time.Millisecond)
	if inventory.ticket <= 0 {
		return false
	}
	inventory.ticket -= 1
	return true
}

func worker(ctx context.Context, users <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for userId := range users {
		select {
		case <-ctx.Done():
			fmt.Printf("Hết giờ, hủy đơn user: %d\n", userId)
		default:
			isSucceed := inventory.BuyTicket()
			if isSucceed {
				fmt.Printf("User %d mua thành công\n", userId)
			} else {
				fmt.Printf("User %d mua thất bại\n", userId)
			}
		}
	}
}
