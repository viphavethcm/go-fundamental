package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Notifier interface {
	Send(ctx context.Context, message string) error
}
type SMS struct {
}

func (sms SMS) Send(ctx context.Context, message string) error {
	select {
	case <-ctx.Done():
		fmt.Println("Gửi SMS failed due to timeout")
		return ctx.Err()
	case <-time.After(1 * time.Second):
		fmt.Println("Gửi SMS: ", message)
		return nil
	}
}

type PushNotification struct {
}

func (noti PushNotification) Send(ctx context.Context, message string) error {
	select {
	case <-ctx.Done():
		fmt.Println("Gửi Push failed due to timeout")
		return ctx.Err()
	case <-time.After(1 * time.Second):
		fmt.Println("Gửi Push: ", message)
		return nil
	}
}

type Job struct {
	ID       int
	Message  string
	Provider Notifier
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err := job.Provider.Send(context, job.Message)
		if err != nil {
			fmt.Printf("Worker %d xử lý Job %d lỗi: %s \n", id, job.ID, err)
		} else {
			fmt.Printf("Worker %d xử lý Job %d thành công\n", id, job.ID)
		}
		cancel()
	}
}

func main() {
	jobs := make(chan Job, 10)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("message %d", i)
		random := rand.Intn(100)
		job := Job{ID: i, Message: message, Provider: &SMS{}}
		if random%2 == 0 {
			job = Job{ID: i, Message: message, Provider: &PushNotification{}}
		}
		jobs <- job
	}
	close(jobs)
	wg.Wait()
}
