package makecmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cryingcatscloud/runway/internal/maker/model"
	"github.com/spf13/cobra"
)

var ModelCmd = &cobra.Command{
	Use:  "model [name]",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string

		if len(args) > 0 {
			name = args[0]
		} else {
			prompt := &survey.Input{
				Message: "Model name:",
			}

			if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
				return err
			}
		}

		if err := model.Create(name); err != nil {
			return err
		}

		fmt.Println("✔ model created:", name)
		fmt.Println()
		fmt.Println("Don't forget to generate ent code")
		fmt.Println()
		fmt.Println("  crew gen ent")
		fmt.Println()

		return nil
	},
}
