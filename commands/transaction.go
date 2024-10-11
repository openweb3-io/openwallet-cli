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

			apiClient := getAPIClientOrExit()
			l, err := apiClient.Transaction.List(cmd.Context(), getTransactionListOptions(cmd))
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

			apiClient := getAPIClientOrExit()
			out, err := apiClient.Transaction.Retrieve(cmd.Context(), transactionID)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	wc.cmd.AddCommand(get)

	// estimateFee
	currencyFlagName := "data-currency"
	fromFlagName := "data-from"
	toFlagName := "data-to"
	feeTypeFlagName := "data-feeType"
	estimateFee := &cobra.Command{
		Use:   "estimate-fee [JSON_PAYLOAD]",
		Short: "Estimate fee",
		Long: `Estimate fee
	
	Example Schema:
	{
		"currency": "string",
		"from": "string",
		"to": "string",
		"feeType": "string"
	  }
	`,
		Args: validators.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			/*
				printer := pretty.NewPrinter(getPrinterOptions())
				var in []byte
				if len(args) > 0 {
					in = []byte(args[0])
				} else {
					var err error
					in, err = utils.ReadStdin()
					printer.CheckErr(err)
				}

				var req wallet.EstimateFeeIn
				if len(in) > 0 {
					err := json.Unmarshal(in, &req)
					printer.CheckErr(err)
				}

				// get flags
				if cmd.Flags().Changed(currencyFlagName) {
					currencyFlag, err := cmd.Flags().GetString(currencyFlagName)
					printer.CheckErr(err)
					req.Currency = currencyFlag
				}
				if cmd.Flags().Changed(fromFlagName) {
					fromFlag, err := cmd.Flags().GetString(fromFlagName)
					printer.CheckErr(err)
					req.From = fromFlag
				}
				if cmd.Flags().Changed(toFlagName) {
					toFlag, err := cmd.Flags().GetString(toFlagName)
					printer.CheckErr(err)
					req.To = toFlag
				}
				if cmd.Flags().Changed(feeTypeFlagName) {
					feeTypeFlag, err := cmd.Flags().GetString(feeTypeFlagName)
					printer.CheckErr(err)
					req.feeType = feeTypeFlag
				}

				apiClient := getAPIClientOrExit()
				apiClient.Transaction.EstimateFee()
				out, err := walletClient.Webhook.Create(cmd.Context(), &req)
				printer.CheckErr(err)
				printer.Print(out)
			*/
		},
	}
	estimateFee.Flags().String(currencyFlagName, "", "")
	estimateFee.Flags().String(fromFlagName, "", "")
	estimateFee.Flags().String(toFlagName, "", "")
	estimateFee.Flags().String(feeTypeFlagName, "", "")
	wc.cmd.AddCommand(estimateFee)

	return wc
}

func addTransactionFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getTransactionListOptions(cmd *cobra.Command) *wallet.ListTransactionOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.ListTransactionOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
