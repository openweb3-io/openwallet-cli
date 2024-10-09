package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type addressCmd struct {
	cmd *cobra.Command
}

func newAddressCmd() *addressCmd {
	wc := &addressCmd{}
	wc.cmd = &cobra.Command{
		Use:     "address",
		Short:   "List, create addresses",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"address"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List deposit addresses",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Address.List(cmd.Context(), getAddressListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addWalletFilterFlags(list)
	wc.cmd.AddCommand(list)

	return wc
}

func addAddressFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

// TODO cursor
func getAddressListOptions(cmd *cobra.Command) *wallet.ListAddressOptions {
	limit, _ := cmd.Flags().GetInt("limit")

	opts := &wallet.ListAddressOptions{
		Limit: limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = cursorFlag
	}

	return opts
}
