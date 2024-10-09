package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type accountCmd struct {
	cmd *cobra.Command
}

func newAccountCmd() *accountCmd {
	wc := &accountCmd{}
	wc.cmd = &cobra.Command{
		Use:     "account",
		Short:   "List, create & modify accounts",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"account"},
	}

	// list
	list := &cobra.Command{
		Use:   "list WALLET_ID",
		Short: "List wallet accounts",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletID := args[0]

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Wallet.ListAccounts(cmd.Context(), walletID, getAccountListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addAccountFilterFlags(list)
	wc.cmd.AddCommand(list)

	return wc
}

func addAccountFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getAccountListOptions(cmd *cobra.Command) *wallet.ListAccountsOptions {
	limit, _ := cmd.Flags().GetInt("limit")

	opts := &wallet.ListAccountsOptions{
		Limit: limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
