package commands

import (
	"fmt"

	"github.com/openweb3-io/openwallet-cli/validators"
	"github.com/openweb3-io/openwallet-cli/version"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	cmd *cobra.Command
}

func newVersionCmd() *versionCmd {
	return &versionCmd{
		cmd: &cobra.Command{
			Use:   "version",
			Args:  validators.NoArgs(),
			Short: "Get the version of the Wallet CLI",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(version.String())
			},
		},
	}
}
