package mod

import (
	"errors"
	"github.com/spf13/cobra"
)

func GetModCommand() *cobra.Command {
	return &cobra.Command{
		Aliases:[]string{"mod"},
		Use: "mod MODNAME",
		Short: "Scrapes one specific mod from CurseForge",
		Long: "Creates a directory and scrapes information about one specific mod from CurseForge and puts info in directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("only specify MODNAME for mod command")
			}

			getMod(args[0])
			return nil
		},

	}
}

func getMod(modName string) {

}