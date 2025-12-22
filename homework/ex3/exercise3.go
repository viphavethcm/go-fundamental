package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	orders := make(chan string, 2)
	menu := []string{"Pepsi", "Coca", "7Up"}
	fmt.Printf("--- Bắt đầu nhận order ---\n")
	var wg sync.WaitGroup
	wg.Add(1)
	go barista(orders, &wg)

	for _, item := range menu {
		orders <- item
		fmt.Printf("Đã nhận đơn %s vào bếp\n", item)
	}
	close(orders)
	wg.Wait()

	fmt.Println("--- Đã nhận hết đơn, nhân viên đi nghỉ ---")
}

func barista(orders <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range orders {
		fmt.Printf("Barista: Preparing %s....\n", item)
		time.Sleep(1 * time.Second)
		fmt.Printf("Barista: Finished %s....\n", item)
	}
}
