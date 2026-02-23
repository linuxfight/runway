package cli

import (
	"fmt"

	"github.com/cryingcatscloud/runway/internal/cli/gen"
	makePkg "github.com/cryingcatscloud/runway/internal/cli/make"
	"github.com/cryingcatscloud/runway/internal/version"
	"github.com/spf13/cobra"
)

var showVersion bool

var rootCmd = &cobra.Command{
	Use:   "crew",
	Short: "crew is the CLI for the runway framework",
	Long: `crew is a developer tool for scaffolding projects,
modules and APIs built on top of the runway framework.`,
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       version.Version(),
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Println(version.Version())
			return nil
		}

		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "print version")

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

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		current := version.Version()

		latest, err := version.CheckWithCache(current)
		if err != nil {
			cmd.Println("Error checking for updates: ", err.Error())
			return
		}
		if latest == "" {
			return
		}

		cmd.Printf(
			"\nA new version of crew is available: %s (current %s)\n"+
				"Update with:\n"+
				"  go install github.com/cryingcatscloud/runway/cmd/crew@latest\n\n",
			latest,
			current,
		)
	}
}
