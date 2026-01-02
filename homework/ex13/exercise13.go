package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	// Main không cần quan tâm wg, việc đó để merge lo
	c1 := fetchPrice("Shopee")
	c2 := fetchPrice("Lazada")
	c3 := fetchPrice("Tiki")

	// Gom 3 dòng chảy thành 1 dòng chảy lớn
	out := merge(c1, c2, c3)

	// Hứng kết quả từ dòng chảy lớn
	for v := range out {
		fmt.Println(v)
	}
}

func fetchPrice(source string) <-chan string {
	out := make(chan string)
	go func() {
		// Giả lập độ trễ ngẫu nhiên để thấy tính song song
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)

		out <- fmt.Sprintf("Giá từ %s: %d", source, rand.IntN(100)+100)

		// QUAN TRỌNG: Gửi xong phải đóng để báo hiệu hết tin
		close(out)
	}()
	return out
}

func merge(cs ...<-chan string) <-chan string {
	// Channel tổng hợp kết quả
	out := make(chan string)
	var wg sync.WaitGroup

	// Hàm worker trung gian: Nhiệm vụ là copy từ input -> output
	output := func(c <-chan string) {
		defer wg.Done() // Báo cáo xong việc khi channel c bị đóng
		for n := range c {
			out <- n
		}
	}

	// Fan-out: Với mỗi channel đầu vào, tạo 1 worker riêng để lắng nghe nó
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c) // Chạy song song, không chờ nhau
	}

	// Monitor: Một goroutine riêng nằm chờ tất cả worker xong việc thì đóng channel tổng
	go func() {
		wg.Wait()
		close(out) // Đóng channel tổng để hàm main thoát vòng lặp
	}()

	return out
}
