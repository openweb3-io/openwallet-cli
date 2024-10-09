package commands

import (
	"encoding/json"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/utils"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type walletCmd struct {
	cmd *cobra.Command
}

func newWalletCmd() *walletCmd {
	wc := &walletCmd{}
	wc.cmd = &cobra.Command{
		Use:     "wallet",
		Short:   "List, create & modify wallets",
		Args:    validators.ExactArgs(1),
		Aliases: []string{"wallet"},
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List wallets",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Wallet.List(cmd.Context(), getWalletListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addWalletFilterFlags(list)
	wc.cmd.AddCommand(list)

	// create
	nameFlagName := "data-name"
	uidFlagName := "data-uid"
	create := &cobra.Command{
		Use:   "create [JSON_PAYLOAD]",
		Short: "Create a new wallet",
		Long: `Create a new wallet

Example Schema:
{
	"name": "wallet-name",
	"uid": "wallet-uid"
  }
`,
		Args: validators.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			var in []byte
			if len(args) > 1 {
				in = []byte(args[1])
			} else {
				var err error
				in, err = utils.ReadStdin()
				printer.CheckErr(err)
			}
			var wallet wallet.CreateWalletIn
			if len(in) > 0 {
				err := json.Unmarshal(in, &wallet)
				printer.CheckErr(err)
			}

			// get flags
			if cmd.Flags().Changed(nameFlagName) {
				nameFlag, err := cmd.Flags().GetString(nameFlagName)
				printer.CheckErr(err)
				wallet.Name = nameFlag
			}

			if cmd.Flags().Changed(uidFlagName) {
				uidFlag, err := cmd.Flags().GetString(uidFlagName)
				printer.CheckErr(err)
				wallet.SetUid(uidFlag)
			}

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Wallet.Create(cmd.Context(), &wallet)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	create.Flags().String(nameFlagName, "", "")
	create.Flags().String(uidFlagName, "", "")
	wc.cmd.AddCommand(create)

	get := &cobra.Command{
		Use:   "get WALLET_ID",
		Short: "get wallet by id",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletID := args[0]

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Wallet.Retrieve(cmd.Context(), walletID)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	wc.cmd.AddCommand(get)

	update := &cobra.Command{
		Use:   "update WALLET_ID [JSON_PAYLOAD]",
		Short: "Update an wallet by id",
		Long: `Update an wallet by id

Example Schema:
{
  "name": "wallet-name",
  "uid": "uid"
}
`,
		Args: validators.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			// parse args
			walletID := args[0]

			var in []byte
			if len(args) > 2 {
				in = []byte(args[2])
			} else {
				var err error
				in, err = utils.ReadStdin()
				printer.CheckErr(err)
			}
			var wallet wallet.UpdateWalletIn
			if len(in) > 0 {
				err := json.Unmarshal(in, &wallet)
				printer.CheckErr(err)
			}

			// get flags
			if cmd.Flags().Changed(nameFlagName) {
				nameFlag, err := cmd.Flags().GetString(nameFlagName)
				printer.CheckErr(err)
				wallet.SetName(nameFlag)
			}

			if cmd.Flags().Changed(uidFlagName) {
				uidFlag, err := cmd.Flags().GetString(uidFlagName)
				printer.CheckErr(err)
				wallet.SetUid(uidFlag)
			}

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Wallet.Update(cmd.Context(), walletID, &wallet)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	update.Flags().String(nameFlagName, "", "")
	update.Flags().String(uidFlagName, "", "")
	wc.cmd.AddCommand(update)

	/*
		delete := &cobra.Command{
			Use:   "delete WALLET_ID",
			Short: "Delete an wallet by id",
			Args:  validators.ExactArgs(2),
			Run: func(cmd *cobra.Command, args []string) {

					printer := pretty.NewPrinter(getPrinterOptions())

					// parse args
					appID := args[0]
					walletID := args[1]

					utils.Confirm(fmt.Sprintf("Are you sure you want to delete the the wallet with id: %s", walletID))

					walletClient := getWalletClientOrExit()
					err := walletClient.Wallet.Delete(cmd.Context(), appID, walletID)
					printer.CheckErr(err)

					fmt.Printf("Wallet \"%s\" Deleted!\n", walletID)

			},
		}
		wc.cmd.AddCommand(delete)
	*/

	return wc
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
