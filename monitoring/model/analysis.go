package model

import (
	"context"
	"sync"
	"time"
)

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}
type Analysis struct {
	Attribute string
	Value     string
}

type ProcessAnalysis struct {
	Name        string
	CPU         float64
	Memory      uint64
	RamPercent  float64
	RunningTime time.Duration
}

var (
	StatsMutex sync.Mutex
	Reports    = map[string]Analysis{}
)
