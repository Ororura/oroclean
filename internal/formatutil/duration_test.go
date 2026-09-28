package formatutil

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		name  string
		value time.Duration
		want  string
	}{
		{
			name:  "microseconds",
			value: 420 * time.Microsecond,
			want:  "420µs",
		},
		{
			name:  "milliseconds",
			value: 125 * time.Millisecond,
			want:  "125ms",
		},
		{
			name:  "seconds",
			value: 2*time.Second + 345*time.Millisecond,
			want:  "2.35s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Duration(tt.value)

			if got != tt.want {
				t.Fatalf(
					"Duration(%v) = %q, want %q",
					tt.value,
					got,
					tt.want,
				)
			}
		})
	}
}
