package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string)
	menu := []string{"Cà phê sữa", "Bạc xỉu", "Cacao"}
	go barista(orders)
	for _, item := range menu {
		orders <- item
	}
	close(orders)
	time.Sleep(3 * time.Second)
}
func barista(orders <-chan string) {
	for item := range orders {
		fmt.Printf("Barista: Preparing %s....\n", item)
		time.Sleep(1 * time.Second)
		fmt.Printf("Barista: Finished %s....\n", item)
	}
}
