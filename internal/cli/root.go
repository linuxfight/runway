package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crew",
	Short: "crew is the CLI for the runway framework",
	Long: `crew is a developer tool for scaffolding projects,
modules and APIs built on top of the runway framework.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(genCmd)
}
