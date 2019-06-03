package pack

import "github.com/forgemods/forge-mod-manager/command/launcher/config"

var VersionJSONMap = map[string]*config.InheritedVersion{}

type PackInfo struct {
	Name             string `yaml:"name"`
	Version          string `yaml:"version"`
	ForgeVersion     string `yaml:"forge_version"`
	MinecraftVersion string `yaml:"minecraft_version"`
	JavaVersion      string `yaml:"java_version"`
	Icon             string `yaml:"icon"`
}

type PackModInfo struct {
	Mod     string `yaml:"mod"`
	Version string `yaml:"version"`
}

type PackFile struct {
	Pack *PackInfo      `yaml:"pack"`
	Mods []*PackModInfo `yaml:"mods"`
}

func (p *PackFile) EachMod(f func(*PackModInfo)) {
	for _, mod := range p.Mods {
		f(mod)
	}
}
