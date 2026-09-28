package scanner

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/Ororura/oroclean/internal/model"
)

type FilesystemScanner struct{}

func NewFilesystemScanner() *FilesystemScanner {
	return &FilesystemScanner{}
}

func (s *FilesystemScanner) Scan(
	ctx context.Context,
	target string,
) (model.ScanResult, error) {
	startedAt := time.Now()

	absoluteTarget, err := filepath.Abs(target)
	if err != nil {
		return model.ScanResult{}, fmt.Errorf(
			"resolve target path: %w",
			err,
		)
	}

	result := model.ScanResult{
		Target:    absoluteTarget,
		StartedAt: startedAt,
	}

	err = filepath.WalkDir(
		absoluteTarget,
		func(
			path string,
			entry fs.DirEntry,
			walkErr error,
		) error {
			if err := ctx.Err(); err != nil {
				return err
			}

			if walkErr != nil {
				if path == absoluteTarget {
					return walkErr
				}

				result.SkippedCount++
				result.Issues = append(
					result.Issues,
					model.ScanIssue{
						Path:  path,
						Error: walkErr.Error(),
					},
				)

				return nil
			}

			if entry.Type()&fs.ModeSymlink != 0 {
				return nil
			}

			if entry.IsDir() {
				result.DirCount++
				return nil
			}

			info, err := entry.Info()
			if err != nil {
				result.SkippedCount++
				result.Issues = append(
					result.Issues,
					model.ScanIssue{
						Path:  path,
						Error: err.Error(),
					},
				)

				return nil
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			result.FileCount++
			result.TotalSize += info.Size()

			return nil
		},
	)
	if err != nil {
		return model.ScanResult{}, fmt.Errorf(
			"scan %q: %w",
			absoluteTarget,
			err,
		)
	}

	result.CompletedAt = time.Now()
	result.Duration = result.CompletedAt.Sub(result.StartedAt)

	return result, nil
}
