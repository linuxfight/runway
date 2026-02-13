package makecmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cryingcatscloud/runway/internal/maker/config"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config [name]",
	Short: "Create a config in app/config",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string

		if len(args) > 0 {
			name = args[0]
		} else {
			prompt := &survey.Input{
				Message: "Config name:",
			}

			if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
				return err
			}
		}

		if err := config.Create(name); err != nil {
			return err
		}

		fmt.Println("✔ config created:", name)
		fmt.Println()
		fmt.Println("Don't forget to register it in app/config/module.go:")
		fmt.Println()
		fmt.Printf("    fx.Provide(New%sCfg),\n", config.Title(name))

		return nil
	},
}
