package scanner

import (
	"context"

	"github.com/Ororura/oroclean/internal/model"
)

type Scanner interface {
	Scan(ctx context.Context, target string) (model.ScanResult, error)
}
