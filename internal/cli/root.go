package cli

import (
	"github.com/cryingcatscloud/runway/internal/cli/gen"
	makePkg "github.com/cryingcatscloud/runway/internal/cli/make"
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
	rootCmd.AddCommand(makeCmd)

	genCmd.AddCommand(gen.EntCmd)
	genCmd.AddCommand(gen.ServerCmd)
	genCmd.AddCommand(gen.OpenAPICmd)
	genCmd.AddCommand(gen.AllCmd)

	makeCmd.AddCommand(makePkg.ConfigCmd)
	makeCmd.AddCommand(makePkg.ModuleCmd)
	makeCmd.AddCommand(makePkg.ModelCmd)
}
