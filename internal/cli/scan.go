package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Ororura/oroclean/internal/formatutil"
	"github.com/Ororura/oroclean/internal/model"
	"github.com/Ororura/oroclean/internal/scanner"
	"github.com/spf13/cobra"
)

func newScanCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "scan <path>",
		Short: "Analyze disk usage for a directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filesystemScanner := scanner.NewFilesystemScanner()

			result, err := filesystemScanner.Scan(
				cmd.Context(),
				args[0],
			)
			if err != nil {
				return err
			}

			if jsonOutput {
				return writeScanJSON(cmd, result)
			}

			writeScanText(cmd, result)

			return nil
		},
	}

	cmd.Flags().BoolVar(
		&jsonOutput,
		"json",
		false,
		"Output scan result as JSON",
	)

	return cmd
}

func writeScanJSON(
	cmd *cobra.Command,
	result model.ScanResult,
) error {
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("encode scan result: %w", err)
	}

	return nil
}

func writeScanText(
	cmd *cobra.Command,
	result model.ScanResult,
) {
	out := cmd.OutOrStdout()

	fmt.Fprintf(out, "Target:      %s\n", result.Target)
	fmt.Fprintf(
		out,
		"Size:        %s\n",
		formatutil.Bytes(result.TotalSize),
	)
	fmt.Fprintf(out, "Files:       %d\n", result.FileCount)
	fmt.Fprintf(out, "Directories: %d\n", result.DirCount)
	fmt.Fprintf(out, "Skipped:     %d\n", result.SkippedCount)
	fmt.Fprintf(
		out,
		"Duration:    %s\n",
		formatutil.Duration(result.Duration),
	)
}
