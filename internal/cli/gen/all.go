package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AllCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all generators",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("→ generating server...")
		if err := ServerCmd.RunE(cmd, args); err != nil {
			return err
		}

		fmt.Println("→ generating openapi...")
		if err := OpenAPICmd.RunE(cmd, args); err != nil {
			return err
		}

		fmt.Println("→ generating ent...")
		if err := EntCmd.RunE(cmd, args); err != nil {
			return err
		}

		fmt.Println("✔ all generators completed")
		return nil
	},
}
