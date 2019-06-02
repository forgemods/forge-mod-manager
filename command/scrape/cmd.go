package scrape

import (
	"github.com/forgemods/forge-mod-manager/command/scrape/all"
	"github.com/forgemods/forge-mod-manager/command/scrape/mod"
	"github.com/spf13/cobra"
)

var scrapeCommand = &cobra.Command{
	Use: "scrape [command]",
	Short: "Forge manager - scrape content from CurseForge",
	Long:  `Tool for scraping mods and mod versions off of the CurseForge website`,
}

func init() {
	// scrape all mods from CurseForge
	scrapeCommand.AddCommand(all.GetAllCommand())

	// scrape one specific mod from CurseForge
	scrapeCommand.AddCommand(mod.GetModCommand())
}


func GetScrapeCommand() *cobra.Command {
	return scrapeCommand
}