package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestScanResultJSON(t *testing.T) {
	startedAt := time.Date(2026, 9, 28, 12, 0, 0, 0, time.UTC)
	completedAt := startedAt.Add(2 * time.Second)

	result := ScanResult{
		Target:      "/tmp/example",
		TotalSize:   1024,
		FileCount:   2,
		DirCount:    1,
		StartedAt:   startedAt,
		CompletedAt: completedAt,
		Duration:    2 * time.Second,
		Items: []Item{
			{
				Path:     "/tmp/example/cache",
				Size:     1024,
				Category: CategoryCache,
				Risk:     RiskSafe,
			},
		},
	}

	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal ScanResult: %v", err)
	}

	var decoded ScanResult

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal ScanResult: %v", err)
	}

	if decoded.Target != result.Target {
		t.Fatalf("Target = %q, want %q", decoded.Target, result.Target)
	}

	if decoded.TotalSize != result.TotalSize {
		t.Fatalf("TotalSize = %d, want %d", decoded.TotalSize, result.TotalSize)
	}

	if len(decoded.Items) != 1 {
		t.Fatalf("len(Items) = %d, want 1", len(decoded.Items))
	}

	if decoded.Items[0].Risk != RiskSafe {
		t.Fatalf(
			"Items[0].Risk = %q, want %q",
			decoded.Items[0].Risk,
			RiskSafe,
		)
	}
}
