package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type rateCmd struct {
	cmd *cobra.Command
}

func newRateCmd() *rateCmd {
	ec := &rateCmd{}
	ec.cmd = &cobra.Command{
		Use:   "rate",
		Short: "query rates",
	}

	estimate := &cobra.Command{
		Use:   "estimate BASE_CURRENCY TO_CURRENCY AMOUNT",
		Short: "estimate",
		Long:  `estimate`,
		Args:  validators.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			baseCurrency := args[0]
			toCurrency := args[1]
			baseAmount := args[2]

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Rate.Estimate(cmd.Context(), &wallet.EstimateOptions{
				BaseCurrency: baseCurrency,
				BaseAmount:   baseAmount,
				ToCurrency:   toCurrency,
			})
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	ec.cmd.AddCommand(estimate)

	get := &cobra.Command{
		Use:   "get BASE_CURRENCY TO_CURRENCY",
		Short: "get",
		Long:  `get`,
		Args:  validators.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			baseCurrency := args[0]
			toCurrency := args[1]

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Rate.GetRates(cmd.Context(), &wallet.GetRatesIn{
				Pairs: []wallet.CurrencyPair{
					{
						BaseCurrency: baseCurrency,
						ToCurrency:   toCurrency,
					},
				},
			})
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	ec.cmd.AddCommand(get)

	return ec
}
