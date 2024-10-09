package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var openableURLs = map[string]string{
	"docs": "https://docs.openweb3.io/",
	"api":  defaultApiUrl,
}

type openCmd struct {
	cmd *cobra.Command
}

func keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func newOpenCmd() *openCmd {
	keys := keys(openableURLs)
	oc := &openCmd{
		cmd: &cobra.Command{
			Use:       fmt.Sprintf("open [%s]", strings.Join(keys, "|")),
			ValidArgs: keys,
			Args:      validators.ExactValidArgs(1),
			Short:     "Quickly open OpenWallet pages in your browser",
			Long: `Quickly open OpenWallet pages in your browser:
docs - opens the OpenWallet documentation
api  - opens the OpenWallet API documentation
			`,
			Run: func(cmd *cobra.Command, args []string) {
				url := openableURLs[args[0]]
				err := open.Run(url)
				if err != nil {
					fmt.Fprintf(os.Stderr, `Failed to open %s in your default browser
To open it manually navigate to:
%s
`, args[0], pretty.MakeTerminalLink(url, url))
					os.Exit(1)
				}
			},
		},
	}
	return oc
}
