package cli

import (
	"fmt"

	"github.com/Ororura/oroclean/internal/formatutil"
	"github.com/Ororura/oroclean/internal/scanner"
	"github.com/spf13/cobra"
)

func newScanCommand() *cobra.Command {
	return &cobra.Command{
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

			out := cmd.OutOrStdout()

			fmt.Fprintf(out, "Target:      %s\n", result.Target)
			fmt.Fprintf(
				out,
				"Size:        %s\n",
				formatutil.Bytes(result.TotalSize),
			)
			fmt.Fprintf(out, "Files:       %d\n", result.FileCount)
			fmt.Fprintf(out, "Directories: %d\n", result.DirCount)
			fmt.Fprintf(
				out,
				"Duration:    %s\n",
				result.Duration,
			)

			return nil
		},
	}
}
