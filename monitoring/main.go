package main

import (
	"context"
	"fmt"
	"fundamental/monitoring/model"
	"fundamental/monitoring/processor"
	"sync"
	"time"

	_ "github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	monitors := []model.Monitor{
		&model.CPUMonitor{},
		&model.MemMonitor{},
		&model.DiskMonitor{},
		&model.NetworkMonitor{},
	}

	var wg sync.WaitGroup
	reportChannel := make(chan model.Analysis)
	for _, monitor := range monitors {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, reportChannel, monitor)
	}
	go func() {
		for report := range reportChannel {
			model.StatsMutex.Lock()
			model.Reports[report.Attribute] = report
			model.StatsMutex.Unlock()
		}
	}()

	printTicker := time.NewTicker(3 * time.Second)
	go func() {
		for range printTicker.C {
			fmt.Println("---- System Status ----")
			for _, report := range model.Reports {
				model.StatsMutex.Lock()
				fmt.Printf("[%s] %s\n", report.Attribute, report.Value)
				model.StatsMutex.Unlock()
			}
			fmt.Println(processor.GetTopProcesses(ctx))
		}
	}()
	time.Sleep(60 * time.Second)
	cancel()
	wg.Wait()
	close(reportChannel)
	printTicker.Stop()
}
