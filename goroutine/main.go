package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d bắt đầu\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc\n", id)
}

func syncWaitGroup() {
	fmt.Println("SyncWaitGroup Example")
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Thời gian thực hiện %s\n", time.Since(start))
}

func unbufferedChannelEx(ch chan int) {
	fmt.Println("Unbuffered Channel Example")
	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
	}()
}

func bufferedChannelEx(ch chan int) {
	fmt.Println("Buffered Channel Example")
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func selectEx() {
	fmt.Println("Select Example")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "Data from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Data from channel 2"
	}()
	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	syncWaitGroup()
	fmt.Println("-----------------------")

	// unbuffered channel
	unbufferedChannel := make(chan int)
	unbufferedChannelEx(unbufferedChannel)
	for value := range unbufferedChannel {
		fmt.Println(value)
	}
	fmt.Println("-----------------------")

	// buffered channel
	bufferedChannel := make(chan int, 3)
	bufferedChannelEx(bufferedChannel)
	for i := 1; i <= 3; i++ {
		fmt.Println(<-bufferedChannel)
	}
	fmt.Println("-----------------------")

	// select example
	selectEx()
	fmt.Println("-----------------------")
}
