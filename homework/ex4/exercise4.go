package main

import "fmt"

func main() {
	chNumbers := make(chan int)
	chResults := make(chan int)
	go generator(chNumbers)
	go square(chNumbers, chResults)
	printer(chResults)
}

func generator(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

func printer(in <-chan int) {
	for num := range in {
		fmt.Println("Kết quả: ", num)
	}
}
