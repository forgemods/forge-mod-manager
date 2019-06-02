package launcher

import (
	"github.com/forgemods/forge-mod-manager/command/launcher/install"
	"github.com/forgemods/forge-mod-manager/command/launcher/versions"
	"github.com/spf13/cobra"
	"log"
)

var launcherCommand = &cobra.Command{
	Use:   "launcher",
	Short: "Forge manager - Config Minecraft launcher",
	Long:  `Tool for configuring and installing Forge versions in Minecraft Launcher`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Launcher.")
	},
}

func init() {
	launcherCommand.AddCommand(versions.GetCommand())

	launcherCommand.AddCommand(install.GetCommand())
}


func GetLauncherCommand() *cobra.Command {
	return launcherCommand
}