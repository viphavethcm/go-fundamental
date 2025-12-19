package model

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemMonitor struct {
}

func (m *MemMonitor) Name() string {
	return "Memory"
}

func (m *MemMonitor) Check(ctx context.Context) string {
	vmStat, error := mem.VirtualMemoryWithContext(ctx)
	if error != nil {
		return "[Memory Monitor] Could not get Memory stats"
	}
	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)
	return value
}
