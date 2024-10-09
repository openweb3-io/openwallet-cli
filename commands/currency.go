package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type currencyCmd struct {
	cmd *cobra.Command
}

func newCurrencyCmd() *currencyCmd {
	wc := &currencyCmd{}
	wc.cmd = &cobra.Command{
		Use:     "currency",
		Short:   "List, get currencies",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"currency"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List currencies",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Currency.List(cmd.Context(), getCurrencyListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addTransactionFilterFlags(list)
	wc.cmd.AddCommand(list)

	return wc
}

func addCurrencyFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getCurrencyListOptions(cmd *cobra.Command) *wallet.CurrencyListOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.CurrencyListOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
