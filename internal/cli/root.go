package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "oroclean",
		Short:        "Safe and transparent macOS cleanup utility",
		SilenceUsage: true,
	}

	return rootCmd
}

func Execute() error {
	return NewRootCommand().Execute()
}
