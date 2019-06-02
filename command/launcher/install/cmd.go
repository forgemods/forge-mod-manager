package install

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/forgemods/forge-mod-manager/command/launcher/common"
	"github.com/forgemods/forge-mod-manager/command/launcher/config"
	"github.com/forgemods/forge-mod-manager/command/launcher/install/pack"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var packFileName string
var minecraftVersion string
var forgeVersion string
var useLatest bool
var server bool

var allModInstances = make(map[string]*Mod)

var installCommand = &cobra.Command{
	Use:   "install [FORGE_PACK_FILE]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Install Forge to Minecraft launcher",
	Long: "Install Forge to Minecraft launcher.\n\n" +
		"Specify either Forge version, Minecraft version or both, to download the package correct for you.\n" +
		"If there is a recommended version, the recommended version will be installed (unless --latest is specified)",
	RunE: func(cmd *cobra.Command, args []string) error {
		var pwdPackFileExist bool

		if _, err := os.Stat("./forgepack.yml"); err == nil {
			pwdPackFileExist = true
		}

		if len(args) == 1 || pwdPackFileExist {
			fmt.Println("Forge pack file found, all other arguments will be ignored.")
			fmt.Println("-----------------------------------------------------------")
			if pwdPackFileExist {
				packFileName = "./forgepack.yml"
			} else {
				packFileName = args[0]
			}
		}

		return doInstall(cmd)
	},
}

func init() {
	installCommand.Flags().StringVar(&minecraftVersion, "mc-version", "", "Minecraft Version, eg. 1.12.2")
	installCommand.Flags().StringVar(&forgeVersion, "forge-version", "", "Forge Version, eg. 14.23.5.2823")
	installCommand.Flags().BoolVar(&useLatest, "latest", false, "Install Latest version of Forge, instead of Recommended")
	installCommand.Flags().BoolVar(&server, "server", false, "Install server version instead of client in launcher")
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func GetCommand() *cobra.Command {
	return installCommand
}

func doInstall(cmd *cobra.Command) error {
	var forgePackFile *pack.PackFile

	if packFileName != "" {
		packFileBody, err := ioutil.ReadFile(packFileName)
		if err != nil {
			return fmt.Errorf("unable to read packfile. Err: %s", err)
		}

		forgePackFile = &pack.PackFile{}
		err = yaml.Unmarshal(packFileBody, forgePackFile)
		if err != nil {
			return fmt.Errorf("unable to de-serialize packfile from yaml format. Err: %s", err)
		}
	} else {
		if minecraftVersion == "" && forgeVersion == "" {
			fmt.Println("Error! No proper arguments were given, please specify at least one of the following:")
			fmt.Println("  [FORGE_PACK_FILE]")
			fmt.Println("  --mc-version string")
			fmt.Println("  --forge-version string")
			fmt.Println("")
			return cmd.Help()
		}
	}

	if forgePackFile != nil {
		if forgeVersion == "" {
			forgeVersion = forgePackFile.Pack.ForgeVersion
		}

		if minecraftVersion == "" {
			minecraftVersion = forgePackFile.Pack.MinecraftVersion
		}
	}

	if minecraftVersion == "" {
		// "25.0" 1.13.2

		// "14.23" 1.12.2
		// "14.22" 1.12.1
		// "14.21" 1.12
		// "14.21" 1.12

		// "13.20" 1.11.2
		// "13.19" 1.11

		// "12.18.0.1" 1.10
		// "12.18.0.2" 1.10.2

		// TODO: Attempt getting minecraft version from forge version
		panic("implement me")
	}

	if forgeVersion == "" {
		// TODO: Get forge version from minecraft version
		panic("implement me")
	}

	if server {
		panic("implement me")
		// TODO: Server: Download Forge, in $PWD
		// downloadForge()

		// TODO: Server: Install minecraft server jar with correct version in $PWD
	} else {
		forgeVersionJsonID := fmt.Sprintf("forge-%s", forgeVersion)
		forgeVersionJsonFileName := fmt.Sprintf("%s.json", forgeVersionJsonID)
		forgeVersionDirectory := path.Join(common.ConfigDir, "versions", forgeVersionJsonID)

		if err := os.MkdirAll(forgeVersionDirectory, 0755); err != nil {
			return fmt.Errorf("unable to create forge version directory, err: %s", err)
		}

		forgeVersionLauncherPath := path.Join(forgeVersionDirectory, forgeVersionJsonFileName)
		if _, err := os.Stat(forgeVersionLauncherPath); os.IsNotExist(err) {
			if forgeVersionJson, err := json.Marshal(pack.VersionJSONMap[forgeVersionJsonID]); err != nil {
				return fmt.Errorf("unable to marshal VersionJSONMap to JSON, err: %s", err)
			} else {
				if err := ioutil.WriteFile(forgeVersionLauncherPath, []byte(forgeVersionJson), 0644); err != nil {
					return fmt.Errorf("unable to write version file for forge, err: %s", err)
				}

				log.Printf("Json file: %s", forgeVersionJson)
				log.Printf("Json ID: %s", forgeVersionJsonID)
				log.Printf("version dir: %s", forgeVersionDirectory)
				log.Printf("Launcher path: %s", forgeVersionLauncherPath)
				return fmt.Errorf("unable to find a suitable forge version json")
			}
		}
	}

	if forgePackFile != nil {
		if !server {
			launcherProfilesFilePath := path.Join(common.ConfigDir, "launcher_profiles.json")
			launcherProfilesJson, err := ioutil.ReadFile(launcherProfilesFilePath)
			if err != nil {
				return err
			}

			var launcherProfiles = &config.ProfileSettings{}

			err = json.Unmarshal(launcherProfilesJson, launcherProfiles)
			if err != nil {
				return err
			}

			gameDirPath := path.Join(common.ConfigDir, forgePackFile.Pack.Name)

			if err := os.MkdirAll(gameDirPath, 0755); err != nil && !os.IsExist(err) {
				return err
			}

			launcherProfiles.Profiles[forgePackFile.Pack.Name] = &config.Profile{
				Name: forgePackFile.Pack.Name,
				Icon: forgePackFile.Pack.Icon,
				LastVersionID: fmt.Sprintf("forge-%s", forgePackFile.Pack.ForgeVersion),
				GameDir: gameDirPath,
				JavaArgs: "-Xmx8G -XX:+UnlockExperimentalVMOptions -XX:+UseG1GC -XX:G1NewSizePercent=20 -XX:G1ReservePercent=20 -XX:MaxGCPauseMillis=50 -XX:G1HeapRegionSize=32M",
			}

			launcherProfilesSave, err := json.MarshalIndent(launcherProfiles, "", "  ")
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(launcherProfilesFilePath, launcherProfilesSave, 0644)
			if err != nil {
				return err
			}

			modDirPath := path.Join(gameDirPath, "mods")

			installMods(forgePackFile.Mods, modDirPath)

			if _, err := os.Stat("config"); err == nil {
				_ = filepath.Walk("config", func(filepath string, info os.FileInfo, err error) error {
					if err == nil && !info.IsDir() {
						log.Printf("Copying config file `%s`", filepath)
						_ = os.MkdirAll(path.Join(gameDirPath, path.Dir(filepath)), 0755)
						_, _ = copyFile(filepath, path.Join(gameDirPath, filepath))
					}

					return err
				})
			}
		} else {
			// TODO: Server install mods in $PWD/mods
		}
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var allMods []string

func addMod(modname string) {
	if !stringInSlice(modname, allMods) {
		allMods = append(allMods, modname)
	}
}

func installMods(mods []*pack.PackModInfo, modDir string) {
	log.Printf("Installing mods into directory `%s`", modDir)

	if err := os.MkdirAll(modDir, 0755); err != nil {
		panic(err)
	}

	total := len(mods)

	for i, mod := range mods {
		log.Printf("Resolving dependencies (%d/%d)", i+1, total)
		addMod(mod.Mod)
		deps := getModDependencies(mod.Mod)

		log.Printf("dependencies: %v", deps)

		if _, ok := allModInstances[mod.Mod]; !ok {
			allModInstances[mod.Mod] = &Mod{name: mod.Mod, Dependencies: deps}
		}

		for _, dep := range deps {
			addMod(dep)
		}
	}

	for i, mod := range mods {
		log.Printf("Downloading mod (%d/%d) and its dependencies for %s", i+1, total, mod.Mod)
		jarFilePath := fmt.Sprintf("%s/%s.jar", modDir, mod.Mod)
		if _, err := os.Stat(jarFilePath); os.IsNotExist(err) {
			err = GetMod(mod, modDir, minecraftVersion)
			if err != nil {
				log.Printf("ERROR: could not get mod properly. err: %s", err)
			}

			for _, dep := range allModInstances[mod.Mod].Dependencies {
				err = GetMod(&pack.PackModInfo{Mod:dep}, modDir, minecraftVersion)
				if err != nil {
					log.Printf("ERROR: could not get mod dependency properly. err: %s", err)
				}
			}
		}

		if _, err := zip.OpenReader(jarFilePath); err != nil {
			log.Printf("WARNING: Jar file seems corrupt `%s`", jarFilePath)
		}
	}
}

func downloadForge() {
	// Servers only
	// TODO: Find correct forge download link
	// TODO: Download forge to $PWD
	// TODO: Check file against SHA1 hash, if not correct retry download.
}