package version

import (
	"fmt"
)

var Version = "source"

func String() string {
	return fmt.Sprintf("openwallet version: %s\n", Version)
}
