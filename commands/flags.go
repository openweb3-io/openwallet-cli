package commands

import (
	"os"
	"runtime"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getPrinterOptions(cmd *cobra.Command) *pretty.PrinterOptions {
	colorFlag := viper.GetString("color")
	color := false
	switch colorFlag {
	case "always":
		color = true
	case "never":
		color = false
	default:
		if runtime.GOOS != "windows" {
			isTTY, _, err := utils.IsTTY(os.Stdout)
			if err == nil {
				// just defaults to false if an error occurs
				color = isTTY
			}
		}
	}

	return &pretty.PrinterOptions{
		Color: color,
	}
}
