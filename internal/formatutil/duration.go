package formatutil

import "time"

func Duration(value time.Duration) string {
	switch {
	case value < time.Millisecond:
		return value.Round(time.Microsecond).String()
	case value < time.Second:
		return value.Round(time.Millisecond).String()
	default:
		return value.Round(10 * time.Millisecond).String()
	}
}
