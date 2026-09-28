package formatutil

import "fmt"

const (
	kib = 1024
	mib = kib * 1024
	gib = mib * 1024
	tib = gib * 1024
)

func Bytes(size int64) string {
	switch {
	case size >= tib:
		return fmt.Sprintf("%.2f TiB", float64(size)/tib)
	case size >= gib:
		return fmt.Sprintf("%.2f GiB", float64(size)/gib)
	case size >= mib:
		return fmt.Sprintf("%.2f MiB", float64(size)/mib)
	case size >= kib:
		return fmt.Sprintf("%.2f KiB", float64(size)/kib)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
