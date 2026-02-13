package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
	scaffold "github.com/cryingcatscloud/runway/internal/maker/project"
	"github.com/spf13/cobra"
)

func init() {
	newCmd.Flags().Bool("no-interactive", false, "Disable interactive mode")
	newCmd.Flags().String("req-infra", "", "Required infrastructure: pg,redis")
	newCmd.Flags().String("req-orm", "", "ORM: ent,no-orm")
	newCmd.Flags().Bool("skip-tidy", false, "Skip running 'go mod tidy' after project creation")

}

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new runway project",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		noInteractive, _ := cmd.Flags().GetBool("no-interactive")

		var opts scaffold.ProjectOptions
		var err error

		if noInteractive {
			opts, err = runNonInteractive(cmd, args)
		} else {
			opts, err = runInteractive(cmd, args)
		}
		if err != nil {
			return err
		}

		if _, err := os.Stat(opts.ProjectName); err == nil {
			return fmt.Errorf("directory %q already exists", opts.ProjectName)
		}

		ctx := scaffold.ProjectContext{
			ProjectName: opts.ProjectName,
			ModulePath:  opts.ProjectName,
			Options:     opts,
		}

		if err := scaffold.CreateProject(filepath.Clean(opts.ProjectName), ctx); err != nil {
			return err
		}

		skipTidy, _ := cmd.Flags().GetBool("skip-tidy")
		if !skipTidy {
			fmt.Println("Running go mod tidy...")
			if err := runGoModTidy(opts.ProjectName); err != nil {
				return err
			}
		}

		fmt.Println("✔ Project created:", opts.ProjectName)
		fmt.Println()
		fmt.Println("Next steps:")
		fmt.Printf("  cd %s\n", opts.ProjectName)
		fmt.Println("  crew gen all")
		fmt.Println("  go run ./cmd/" + opts.ProjectName)

		return nil
	},
}

func runInteractive(cmd *cobra.Command, args []string) (scaffold.ProjectOptions, error) {
	opts := scaffold.ProjectOptions{}

	// project name
	if len(args) > 0 {
		opts.ProjectName = args[0]
	} else {
		if err := survey.AskOne(&survey.Input{
			Message: "Project name:",
		}, &opts.ProjectName); err != nil {
			return opts, err
		}
	}

	// infra
	var infra []string
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "Select required infrastructure:",
		Options: []string{"Postgres", "Redis"},
	}, &infra); err != nil {
		return opts, err
	}

	for _, v := range infra {
		switch v {
		case "Postgres":
			opts.Infra.Postgres = true
		case "Redis":
			opts.Infra.Redis = true
		}
	}

	if opts.Infra.Postgres {
		var ormChoice string
		if err := survey.AskOne(&survey.Select{
			Message: "Ok! Now select an ORM for Postgres (use ↑↓ + Enter):",
			Options: []string{
				"Ent (recommended)",
				"No ORM (raw SQL)",
			},
			Default: "Ent (recommended)",
		}, &ormChoice); err != nil {
			return opts, err
		}

		switch ormChoice {
		case "Ent (recommended)":
			opts.ORM.Ent = true
		case "No ORM (raw SQL)":
			opts.ORM.Ent = false
		}
	}

	return opts, nil
}

func runNonInteractive(cmd *cobra.Command, args []string) (scaffold.ProjectOptions, error) {
	var opts scaffold.ProjectOptions

	if len(args) != 1 {
		return opts, fmt.Errorf("project name is required")
	}
	opts.ProjectName = args[0]

	infra, _ := cmd.Flags().GetString("req-infra")
	orm, _ := cmd.Flags().GetString("req-orm")

	for _, v := range strings.Split(infra, ",") {
		switch v {
		case "pg":
			opts.Infra.Postgres = true
		case "redis":
			opts.Infra.Redis = true
		case "":
		default:
			return opts, fmt.Errorf("unknown infra: %s", v)
		}
	}

	if orm == "ent" {
		if !opts.Infra.Postgres {
			return opts, fmt.Errorf("ent requires postgres")
		}
		opts.ORM.Ent = true
	}

	return opts, nil
}

func runGoModTidy(projectDir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
