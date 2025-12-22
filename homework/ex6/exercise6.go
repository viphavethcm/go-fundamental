package main

import (
	"fmt"
	"time"
)

func main() {
	cacheChannel := make(chan string)
	dbChannel := make(chan string)
	go queryCache(cacheChannel)
	go queryDB(dbChannel)
	select {
	case <-cacheChannel:
		fmt.Println("cache data")
	case <-dbChannel:
		fmt.Println("db data")
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("timeout")
	}
}

func queryCache(ch chan string) {
	fmt.Println("Start querying from cache")
	time.Sleep(3 * time.Second)
	fmt.Println("Start querying from cache")
}

func queryDB(ch chan string) {
	fmt.Println("Start querying from DB")
	time.Sleep(2 * time.Second)
	fmt.Println("Start querying from DB")
}
