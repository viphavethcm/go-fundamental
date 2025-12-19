package processor

import (
	"context"
	"fmt"
	"fundamental/monitoring/model"
	"sort"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, channel chan<- model.Analysis, m model.Monitor) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cpuPercent := m.Check(ctx)
			channel <- model.Analysis{Attribute: m.Name(), Value: cpuPercent}
		}
	}
}

func GetTopProcesses(ctx context.Context) string {
	output := ""
	vmStat, _ := mem.VirtualMemoryWithContext(ctx)
	totalMemory := vmStat.Total
	processes, error := process.ProcessesWithContext(ctx)
	if error != nil {
		return "[Processes Monitor] Could not get top processes stats"
	}
	var wg sync.WaitGroup
	var cpuList, memList []model.ProcessAnalysis
	processChannel := make(chan model.ProcessAnalysis, len(processes))

	for _, p := range processes {
		wg.Add(1)
		go func(process *process.Process) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				name, error := process.NameWithContext(ctx)
				if error != nil {
					return
				}
				cpuPercent, error := process.CPUPercentWithContext(ctx)
				if error != nil {
					return
				}
				memInfo, error := process.MemoryInfoWithContext(ctx)
				if error != nil {
					return
				}
				ramPercent := (float64(memInfo.RSS) / float64(totalMemory)) * 100
				if cpuPercent <= 1 || ramPercent <= 1 {
					return
				}
				createTime, error := process.CreateTimeWithContext(ctx)
				if error != nil {
					return
				}
				runningTime := time.Since(time.Unix(createTime/1000, 0))
				processAnalysis := model.ProcessAnalysis{
					Name:        name,
					CPU:         cpuPercent,
					Memory:      memInfo.RSS,
					RamPercent:  ramPercent,
					RunningTime: runningTime,
				}
				processChannel <- processAnalysis
			}
		}(p)
	}
	go func() {
		wg.Wait()
		close(processChannel)
	}()
	for process := range processChannel {
		if process.CPU > 1 {
			cpuList = append(cpuList, process)
		}
		if process.Memory > 1 {
			memList = append(memList, process)
		}
	}
	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPU > cpuList[j].CPU
	})
	sort.Slice(memList, func(i, j int) bool {
		return memList[i].RamPercent > memList[j].RamPercent
	})
	output += "Top 5 CPU consuming processes\n"
	for _, processCPU := range cpuList {
		output += fmt.Sprintf("Process Name: %s, CPU: %.2f%% - RAM: %.2f MB(%.2f%%) - Running: %s \n",
			processCPU.Name,
			processCPU.CPU,
			float64(processCPU.Memory)/1024/1024,
			processCPU.RamPercent,
			processCPU.RunningTime,
		)
	}

	output += "Top 5 Mem consuming processes\n"
	for _, processMem := range memList {
		output += fmt.Sprintf("Process Name: %s, CPU: %.2f%% - RAM: %.2f MB(%.2f%%) - Running: %s \n",
			processMem.Name,
			processMem.CPU,
			float64(processMem.Memory)/1024/1024,
			processMem.RamPercent,
			processMem.RunningTime,
		)
	}

	return output
}
