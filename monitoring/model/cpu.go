package model

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func (m *CPUMonitor) Name() string {
	return "CPU"
}

type CPUMonitor struct {
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	cpuStat, error := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if error != nil && len(cpuStat) == 0 {
		return "[CPU Monitor] Could not get CPU stats"
	}
	value := fmt.Sprintf("%.2f%%", cpuStat[0])
	return value
}
