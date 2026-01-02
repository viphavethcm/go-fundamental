package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// 1. Tạo Buffered Queue (Mô phỏng Burst Traffic tốt hơn)
	// Để Generator bắn cái vèo là xong, không bị block chờ worker
	queue := make(chan string, 20)
	var wg sync.WaitGroup

	// 2. Chạy Generator (Bắn xong đóng queue luôn)
	go generator(queue)

	// 3. Chạy Worker (Chạy ngầm để Main còn làm việc khác)
	wg.Add(1)
	go worker(queue, &wg)

	// 4. Main mô phỏng Admin chờ 3 giây
	// Đây mới là logic đếm giờ chính xác, không bị reset
	time.Sleep(3 * time.Second)
	fmt.Println("--- ⚠️  ADMIN: Nhận tín hiệu tắt... Đang xả hàng tồn... ---")

	// 5. Graceful Shutdown
	// Lúc này Generator đã đóng queue rồi.
	// Worker sẽ tự động chạy nốt các email còn lại trong queue cho đến khi hết.
	wg.Wait()
	fmt.Println("--- ✅ Server Stopped Gracefully ---")
}

func generator(queue chan<- string) {
	fmt.Println("Generator: Đang nạp 20 email vào hàng đợi...")
	for i := 1; i <= 20; i++ {
		queue <- "email " + strconv.Itoa(i) + "."
	}
	close(queue) // Nạp xong đóng ngay
	fmt.Println("Generator: Đã nạp xong!")
}

func worker(queue <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Khởi tạo Ticker bên ngoài vòng lặp
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Vòng lặp này sẽ tự thoát khi queue bị close VÀ đã lấy hết dữ liệu
	for message := range queue {
		// Chặn tốc độ xử lý lại
		<-ticker.C
		fmt.Println("Worker: Sending message:", message)
	}
	// Khi chạy xuống đây nghĩa là Queue đã rỗng và đã đóng -> An toàn để nghỉ
}
