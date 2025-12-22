package main

import (
	"fmt"
	"time"
)

func main() {
	transactions := make(chan int)
	go func() {
		transactions <- 100
		transactions <- 0
		transactions <- 500
		time.Sleep(1 * time.Second)
		close(transactions)
	}()
	for {
		amount, ok := <-transactions
		if ok {
			fmt.Println("amount: ", amount)
		} else {
			fmt.Println("Closed")
			break
		}
	}
}
