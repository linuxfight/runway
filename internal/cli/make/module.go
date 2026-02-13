package makecmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cryingcatscloud/runway/internal/maker/module"
	"github.com/spf13/cobra"
)

var ModuleCmd = &cobra.Command{
	Use:  "module [name]",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string

		if len(args) > 0 {
			name = args[0]
		} else {
			prompt := &survey.Input{
				Message: "Module name:",
			}

			if err := survey.AskOne(prompt, &name, survey.WithValidator(survey.Required)); err != nil {
				return err
			}
		}

		if err := module.Create(name); err != nil {
			return err
		}

		fmt.Println("✔ module created:", name)
		fmt.Println()
		fmt.Println("Don't forget to register it in internal/modules/module.go:")
		fmt.Println()
		fmt.Printf(`    import "internal/modules/%s"

    fx.Options(
        %s.Module,
    )
`, name, name)

		return nil
	},
}
