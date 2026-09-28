package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
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
