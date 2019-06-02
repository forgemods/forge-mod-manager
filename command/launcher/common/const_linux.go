package common

import (
	"fmt"
	"os"
)

var ConfigDir = fmt.Sprintf("%s/.minecraft", os.Getenv("HOME"))