package model

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/net"
)

func (m *NetworkMonitor) Name() string {
	return "Network"
}

type NetworkMonitor struct {
}

func (m *NetworkMonitor) Check(ctx context.Context) string {
	netStat, error := net.IOCountersWithContext(ctx, false)
	if error != nil && len(netStat) == 0 {
		return "[Network Monitor] Could not get Network stats"
	}
	value := fmt.Sprintf("Send: %d KB, Received: %d KB", netStat[0].BytesSent/1024, netStat[0].BytesRecv/1024)
	return value
}
