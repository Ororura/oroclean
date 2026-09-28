package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Ororura/oroclean/internal/model"
)

func TestScanCommand(t *testing.T) {
	root := t.TempDir()

	file := filepath.Join(root, "example.txt")

	if err := os.WriteFile(
		file,
		[]byte("oroclean"),
		0o600,
	); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cmd := NewRootCommand()

	var output bytes.Buffer

	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{
		"scan",
		root,
	})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute scan command: %v", err)
	}

	result := output.String()

	expected := []string{
		"Target:",
		"Size:",
		"Files:",
		"Directories:",
		"Duration:",
	}

	for _, value := range expected {
		if !strings.Contains(result, value) {
			t.Fatalf(
				"output does not contain %q:\n%s",
				value,
				result,
			)
		}
	}
}

func TestScanCommandJSON(t *testing.T) {
	root := t.TempDir()

	file := filepath.Join(root, "example.txt")

	if err := os.WriteFile(
		file,
		[]byte("oroclean"),
		0o600,
	); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cmd := NewRootCommand()

	var output bytes.Buffer

	cmd.SetOut(&output)
	cmd.SetErr(&output)
	cmd.SetArgs([]string{
		"scan",
		root,
		"--json",
	})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute scan command: %v", err)
	}

	var result model.ScanResult

	if err := json.Unmarshal(
		output.Bytes(),
		&result,
	); err != nil {
		t.Fatalf(
			"decode JSON output: %v\n%s",
			err,
			output.String(),
		)
	}

	if result.Target == "" {
		t.Fatal("Target must not be empty")
	}

	if result.FileCount != 1 {
		t.Fatalf(
			"FileCount = %d, want 1",
			result.FileCount,
		)
	}
}
