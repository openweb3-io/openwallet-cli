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

			apiClient := getAPIClientOrExit()
			l, err := apiClient.Currency.List(cmd.Context(), getCurrencyListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addCurrencyFilterFlags(list)
	wc.cmd.AddCommand(list)

	// get
	get := &cobra.Command{
		Use:   "get CURRENCY_CODE",
		Short: "Get an currency by code",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			code := args[0]

			apiClient := getAPIClientOrExit()
			out, err := apiClient.Currency.FindByCode(cmd.Context(), code)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	wc.cmd.AddCommand(get)

	return wc
}

func addCurrencyFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getCurrencyListOptions(cmd *cobra.Command) *wallet.ListCurrencyOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.ListCurrencyOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
