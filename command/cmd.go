package command

import (
	"github.com/forgemods/forge-mod-manager/command/launcher/install"
	"github.com/forgemods/forge-mod-manager/command/launcher/versions"
	"github.com/forgemods/forge-mod-manager/command/scrape"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fm [OPTIONS] COMMAND [ARG...]",
	Short: "Forge Manager",
	Long:  `Tool for managing Minecraft Forge and its mods`,
	RunE:  install.GetCommand().RunE,
}

func decorateCommands() {
	// Config and install forge versions in Minecraft Launcher
	//rootCmd.AddCommand(launcher.GetLauncherCommand())

	rootCmd.AddCommand(install.GetCommand())
	rootCmd.AddCommand(versions.GetCommand())

	// Scrape mod contents from CurseForge website
	rootCmd.AddCommand(scrape.GetScrapeCommand())
}

func Execute() {
	decorateCommands()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
