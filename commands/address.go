package commands

import (
	"fmt"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/skip2/go-qrcode"
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

			apiClient := getAPIClientOrExit()
			l, err := apiClient.Address.List(cmd.Context(), getAddressListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addAddressFilterFlags(list)
	wc.cmd.AddCommand(list)

	// get
	get := &cobra.Command{
		Use:   "get",
		Short: "Get an deposit address by network",
		Args:  validators.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletID := args[0]
			currency := args[1]
			network := args[2]

			apiClient := getAPIClientOrExit()
			out, err := apiClient.Address.GetDepositAddress(cmd.Context(), walletID, &wallet.GetDepositAddressOptions{
				Currency: currency,
				Network:  &network,
			})
			printer.CheckErr(err)

			printer.Print(out)

			qrCode, err := qrcode.New(out.Address, qrcode.Medium)
			printer.CheckErr(err)

			// 打印二维码到终端
			fmt.Println(qrCode.ToSmallString(true))
		},
	}
	wc.cmd.AddCommand(get)

	return wc
}

func addAddressFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("cursor", "c", "", "cursor for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

// TODO cursor
func getAddressListOptions(cmd *cobra.Command) *wallet.ListAddressOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.ListAddressOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("cursor") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
