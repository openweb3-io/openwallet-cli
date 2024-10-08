package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type walletCmd struct {
	cmd *cobra.Command
}

func newWalletCmd() *walletCmd {
	ac := &walletCmd{}
	ac.cmd = &cobra.Command{
		Use:     "wallet",
		Short:   "List, create & modify wallets",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"wallet"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List current applications",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions(cmd))
			appID := args[0]

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Wallet.List(cmd.Context(), appID, getWalletListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addWalletFilterFlags(list)
	ac.cmd.AddCommand(list)

	return ac
}

func addWalletFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getWalletListOptions(cmd *cobra.Command) *wallet.ListWalletOptions {
	limit, _ := cmd.Flags().GetInt("limit")

	opts := &wallet.ListWalletOptions{
		Limit: limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
