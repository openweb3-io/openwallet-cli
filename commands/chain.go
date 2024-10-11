package commands

import (
	"github.com/openweb3-io/openwallet-cli/validators"
	// wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type chainCmd struct {
	cmd *cobra.Command
}

func newChainCmd() *chainCmd {
	wc := &chainCmd{}
	wc.cmd = &cobra.Command{
		Use:     "chain",
		Short:   "List, get chains",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"chain"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List chains",
		Run: func(cmd *cobra.Command, args []string) {
			/*
				printer := pretty.NewPrinter(getPrinterOptions())

				apiClient := getAPIClientOrExit()
				l, err := apiClient.Chain.List(cmd.Context(), getCurrencyListOptions(cmd))
				printer.CheckErr(err)

				printer.Print(l)
			*/
		},
	}
	addChainFilterFlags(list)
	wc.cmd.AddCommand(list)

	// list
	listEnabled := &cobra.Command{
		Use:   "list-enabled",
		Short: "List enabled chains",
		Run: func(cmd *cobra.Command, args []string) {
			/*
				printer := pretty.NewPrinter(getPrinterOptions())

				apiClient := getAPIClientOrExit()
				l, err := apiClient.Chain.List(cmd.Context(), getCurrencyListOptions(cmd))
				printer.CheckErr(err)

				printer.Print(l)
			*/
		},
	}
	addEnabledChainFilterFlags(listEnabled)
	wc.cmd.AddCommand(listEnabled)

	// get
	get := &cobra.Command{
		Use:   "get CHAIN_ID",
		Short: "Get an chain by id",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			/*
				printer := pretty.NewPrinter(getPrinterOptions())

				chainId := args[0]

				apiClient := getAPIClientOrExit()
				out, err := walletClient.Chain.Get(cmd.Context(), chainId)
				printer.CheckErr(err)
				printer.Print(out)
			*/
		},
	}
	wc.cmd.AddCommand(get)

	return wc
}

func addChainFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func addEnabledChainFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

/*
func getChainListOptions(cmd *cobra.Command) *wallet.ListChainOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.ListChainOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
*/
