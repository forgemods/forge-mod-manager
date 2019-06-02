package versions

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

const forgeVersionIndexUrl = "https://files.minecraftforge.net/maven/net/minecraftforge/forge/index_%s.html"

var minecraftVersion string

var versionsCommand = &cobra.Command{
	Use: "versions MINECRAFT_VERSION",
	Short: "List available Forge versions by Minecraft Version",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		minecraftVersion = args[0]
		return doListVersions()
	},
}

func GetCommand() *cobra.Command {
	return versionsCommand
}

func doListVersions() error {
	response, err := http.Get(fmt.Sprintf(forgeVersionIndexUrl, minecraftVersion))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	htmlDocument, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return err
	}

	var recommendedVersion = ""

	htmlDocument.Find(".download-list .download-version").Each(func(i int, selection *goquery.Selection) {
		var recommended = false

		if selection.Find("i.promo-recommended").Length() > 0 {
			recommended = true
		}
		selection.Children().Remove() // eg. tooltip etc

		version := strings.Trim(selection.Text(), " \t\n\r")
		fmt.Println(version)

		if recommended {
			recommendedVersion = version
		}
	})

	if recommendedVersion != "" {
		fmt.Printf("-----\nRecommended version is: %s\n", recommendedVersion)
	}


	return nil
}