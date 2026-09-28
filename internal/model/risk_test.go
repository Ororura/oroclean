package model

import "testing"

func TestRiskLevelValid(t *testing.T) {
	tests := []struct {
		name  string
		risk  RiskLevel
		valid bool
	}{
		{
			name:  "safe",
			risk:  RiskSafe,
			valid: true,
		},
		{
			name:  "review",
			risk:  RiskReview,
			valid: true,
		},
		{
			name:  "manual",
			risk:  RiskManual,
			valid: true,
		},
		{
			name:  "unknown",
			risk:  RiskLevel("unknown"),
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.risk.Valid(); got != tt.valid {
				t.Fatalf(
					"RiskLevel(%q).Valid() = %v, want %v",
					tt.risk,
					got,
					tt.valid,
				)
			}
		})
	}
}
