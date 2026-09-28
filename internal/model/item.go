package model

import "time"

type Item struct {
	Path        string    `json:"path"`
	Size        int64     `json:"size"`
	Category    Category  `json:"category"`
	Risk        RiskLevel `json:"risk"`
	Description string    `json:"description,omitempty"`
	ModifiedAt  time.Time `json:"modifiedAt,omitempty"`
}
