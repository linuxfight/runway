package cli

import "github.com/spf13/cobra"

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Scaffold helpers (module, config, model)",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return ensureProjectRoot()
	},
}
