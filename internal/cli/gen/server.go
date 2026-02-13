package gen

import (
	"github.com/cryingcatscloud/runway/internal/generator/server"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Generate HTTP server adapters for all modules",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.GenerateAll()
	},
}
