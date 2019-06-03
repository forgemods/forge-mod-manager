package common

import "github.com/kirsle/configdir"

// macOS Minecraft Config directory ~/Library/Application Support/minecraft
var ConfigDir = configdir.LocalConfig("minecraft")