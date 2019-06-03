package common

import "github.com/kirsle/configdir"

// Windows Minecraft Config directory %APPDATA%\.minecraft
var ConfigDir = configdir.LocalConfig(".minecraft")
