package model

type ScanIssue struct {
	Path  string `json:"path"`
	Error string `json:"error"`
}
