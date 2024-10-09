package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type transactionCmd struct {
	cmd *cobra.Command
}

func newTransactionCmd() *transactionCmd {
	wc := &transactionCmd{}
	wc.cmd = &cobra.Command{
		Use:     "transaction",
		Short:   "List, create transactions",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"transaction"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List transactions",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Transaction.List(cmd.Context(), getTransactionListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addTransactionFilterFlags(list)
	wc.cmd.AddCommand(list)

	get := &cobra.Command{
		Use:   "get TRANSACTION_ID",
		Short: "get transaction by id",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			transactionID := args[0]

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Transaction.Retrieve(cmd.Context(), transactionID)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	wc.cmd.AddCommand(get)

	return wc
}

func addTransactionFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getTransactionListOptions(cmd *cobra.Command) *wallet.ListTransactionOptions {
	limit, _ := cmd.Flags().GetInt("limit")

	opts := &wallet.ListTransactionOptions{
		Limit: limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
