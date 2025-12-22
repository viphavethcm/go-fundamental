package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	files := []string{"video.mp4", "image.jpg", "profile.pdf"}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			downloadFile(file, &wg)
		}(file)
	}
	wg.Wait()
}

func downloadFile(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start downloading: ", filename)
	time.Sleep(time.Duration(rand.IntN(3)) * time.Second)
	fmt.Println("Finished downloading: ", filename)
}
