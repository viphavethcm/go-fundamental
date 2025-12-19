package model

import (
	"context"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"
)

func (m *DiskMonitor) Name() string {
	return "Disk"
}

type DiskMonitor struct {
}

func (m *DiskMonitor) Check(ctx context.Context) string {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	diskStat, error := disk.UsageWithContext(ctx, path)
	if error != nil {
		return "[Disk Monitor] Could not get Disk stats"
	}
	value := fmt.Sprintf("%.2f%%", diskStat.UsedPercent)
	return value
}
