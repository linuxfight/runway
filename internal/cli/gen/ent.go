package gen

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var EntCmd = &cobra.Command{
	Use:   "ent",
	Short: "Run go generate for ent schemas",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := "ent"

		c := exec.Command("go", "generate", "./"+dir+"/...")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin

		return c.Run()
	},
}
