package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	jobs := make(chan *Job)
	results := make(chan *Result)
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go generator(jobs, ctx)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, ctx, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	succeeded := 0
	failed := 0
	for result := range results {
		if result != nil && result.Error == nil {
			succeeded += 1
		} else {
			failed += 1
		}
	}
	fmt.Println("Tổng số job thành công: ", succeeded)
	fmt.Println("Tổng số job thất bại: ", failed)
	fmt.Println("Tổng số job timeout: ", 20-(succeeded+failed))
}

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobId  string
	Status string
	Error  error
}

func generator(jobs chan<- *Job, ctx context.Context) {
	for i := 1; i <= 20; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			job := &Job{ID: i, URL: "url_" + strconv.Itoa(i)}
			jobs <- job
		}
	}
	close(jobs)
}

func worker(jobs <-chan *Job, ctx context.Context, results chan<- *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		var err error
		err = nil
		status := "Success"
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			if job.URL == "url_10" {
				err = errors.New("404 Not Found")
				status = "Failed"
			}
			results <- &Result{JobId: strconv.Itoa(job.ID), Status: status, Error: err}
		}
	}
}
