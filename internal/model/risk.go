package model

type RiskLevel string

const (
	RiskSafe   RiskLevel = "safe"
	RiskReview RiskLevel = "review"
	RiskManual RiskLevel = "manual"
)

func (r RiskLevel) Valid() bool {
	switch r {
	case RiskSafe, RiskReview, RiskManual:
		return true
	default:
		return false
	}
}
