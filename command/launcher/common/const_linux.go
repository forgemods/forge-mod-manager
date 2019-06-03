package common

import (
	"fmt"
	"os"
)

// Linux Minecraft Config directory ~/.minecraft
var ConfigDir = fmt.Sprintf("%s/.minecraft", os.Getenv("HOME"))