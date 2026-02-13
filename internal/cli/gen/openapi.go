package gen

import (
	"github.com/cryingcatscloud/runway/internal/generator/openapi"
	"github.com/spf13/cobra"
)

var OpenAPICmd = &cobra.Command{
	Use:   "openapi",
	Short: "Generate openapi.yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return openapi.Generate(openapi.Config{
			Title:   "API",
			Version: "1.0.0",
			OutFile: "openapi.yaml",
		})
	},
}
