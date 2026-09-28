package model

import "time"

type ScanResult struct {
	Target      string        `json:"target"`
	TotalSize   int64         `json:"totalSize"`
	FileCount   int64         `json:"fileCount"`
	DirCount    int64         `json:"dirCount"`
	Items       []Item        `json:"items,omitempty"`
	StartedAt   time.Time     `json:"startedAt"`
	CompletedAt time.Time     `json:"completedAt"`
	Duration    time.Duration `json:"duration"`
}
