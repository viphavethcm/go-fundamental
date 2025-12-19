package main

import (
	"context"
	"fmt"
	"time"
)

func cookCom(ctx context.Context, channelPho chan<- string) {
	fmt.Println("Bat dau nau com")
	select {
	case <-time.After(100 * time.Millisecond):
		channelPho <- "Com da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau com")
		return
	}
}

func cookPho(ctx context.Context, channelPho chan<- string) {
	fmt.Println("Bat dau nau pho")
	select {
	case <-time.After(1 * time.Second):
		channelPho <- "Pho da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau pho")
		return
	}
}

func cookChao(ctx context.Context, channelChao chan<- string) {
	fmt.Println("Bat dau nau chao")
	select {
	case <-time.After(3 * time.Second):
		channelChao <- "chao da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau chao")
		return
	}
}

func employee(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			name := ctx.Value("name")
			fmt.Println("Username: ", name)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func main() {
	channelPho := make(chan string)
	channelChao := make(chan string)
	channelCom := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go cookPho(ctx, channelPho)
	go cookChao(ctx, channelChao)
	go cookCom(ctx, channelCom)
	//ctx = context.WithValue(ctx, "name", "duy")
	//go employee(ctx)
	for i := 1; i <= 3; i++ {
		select {
		case pho := <-channelPho:
			fmt.Println("Nhan duoc: ", pho)
		case chao := <-channelChao:
			fmt.Println("Nhan dc: ", chao)
		case com := <-channelCom:
			fmt.Println("Nhan dc: ", com)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}
