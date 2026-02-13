package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Code generators",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return ensureProjectRoot()
	},
}

func ensureProjectRoot() error {
	missing := func(p string) error {
		return fmt.Errorf("run this command from project root: '%s' not found", p)
	}

	// go.mod
	if _, err := os.Stat("go.mod"); err != nil {
		return missing("go.mod")
	}

	// internal/
	if st, err := os.Stat("internal"); err != nil || !st.IsDir() {
		return missing("internal/")
	}

	// internal/modules/
	modulesPath := filepath.Join("internal", "modules")
	if st, err := os.Stat(modulesPath); err != nil || !st.IsDir() {
		return missing("internal/modules/")
	}

	return nil
}
