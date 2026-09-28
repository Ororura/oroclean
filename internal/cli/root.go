package cli

import (
	"context"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "oroclean",
		Short:        "Safe and transparent macOS cleanup utility",
		SilenceUsage: true,
	}

	rootCmd.AddCommand(
		newVersionCommand(),
		newScanCommand(),
	)

	return rootCmd
}

func Execute() error {
	return ExecuteContext(context.Background())
}

func ExecuteContext(ctx context.Context) error {
	return NewRootCommand().ExecuteContext(ctx)
}
