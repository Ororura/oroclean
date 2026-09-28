package scanner

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFilesystemScanner(t *testing.T) {
	root := t.TempDir()

	firstFile := filepath.Join(root, "first.txt")
	if err := os.WriteFile(
		firstFile,
		[]byte("hello"),
		0o600,
	); err != nil {
		t.Fatalf("write first file: %v", err)
	}

	nested := filepath.Join(root, "nested")
	if err := os.Mkdir(nested, 0o700); err != nil {
		t.Fatalf("create nested directory: %v", err)
	}

	secondFile := filepath.Join(nested, "second.txt")
	if err := os.WriteFile(
		secondFile,
		[]byte("world!"),
		0o600,
	); err != nil {
		t.Fatalf("write second file: %v", err)
	}

	scanner := NewFilesystemScanner()

	result, err := scanner.Scan(context.Background(), root)
	if err != nil {
		t.Fatalf("scan directory: %v", err)
	}

	if result.FileCount != 2 {
		t.Fatalf(
			"FileCount = %d, want 2",
			result.FileCount,
		)
	}

	if result.DirCount != 2 {
		t.Fatalf(
			"DirCount = %d, want 2",
			result.DirCount,
		)
	}

	const expectedSize int64 = 11

	if result.TotalSize != expectedSize {
		t.Fatalf(
			"TotalSize = %d, want %d",
			result.TotalSize,
			expectedSize,
		)
	}

	if result.Target == "" {
		t.Fatal("Target must not be empty")
	}

	if result.Duration < 0 {
		t.Fatalf(
			"Duration = %v, want >= 0",
			result.Duration,
		)
	}
}

func TestFilesystemScannerDoesNotFollowSymlinks(t *testing.T) {
	root := t.TempDir()
	external := t.TempDir()

	externalFile := filepath.Join(external, "large.txt")
	if err := os.WriteFile(
		externalFile,
		[]byte("this file must not be counted"),
		0o600,
	); err != nil {
		t.Fatalf("write external file: %v", err)
	}

	link := filepath.Join(root, "external")

	if err := os.Symlink(external, link); err != nil {
		t.Fatalf("create symlink: %v", err)
	}

	scanner := NewFilesystemScanner()

	result, err := scanner.Scan(context.Background(), root)
	if err != nil {
		t.Fatalf("scan directory: %v", err)
	}

	if result.FileCount != 0 {
		t.Fatalf(
			"FileCount = %d, want 0",
			result.FileCount,
		)
	}

	if result.TotalSize != 0 {
		t.Fatalf(
			"TotalSize = %d, want 0",
			result.TotalSize,
		)
	}
}
