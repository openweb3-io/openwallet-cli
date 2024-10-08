package commands

import (
	"encoding/json"
	"fmt"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/utils"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type endpointCmd struct {
	cmd *cobra.Command
}

func newEndpointCmd() *endpointCmd {
	ec := &endpointCmd{}
	ec.cmd = &cobra.Command{
		Use:   "endpoint",
		Short: "List, create & modify endpoints",
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List current endpoints",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			walletClient := getWalletClientOrExit()
			l, err := walletClient.Webhook.List(cmd.Context(), getEndpointListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addEndpointFilterFlags(list)
	ec.cmd.AddCommand(list)

	// create
	urlFlagName := "data-url"
	versionFlagName := "data-version"
	eventTypesFlagName := "data-eventTypes"
	disabledFlagName := "data-disabled"
	create := &cobra.Command{
		Use:   "create [JSON_PAYLOAD]",
		Short: "Create a new endpoint",
		Long: `Create a new endpoint

Example Schema:
{
	"url": "string",
	"version": 1,
	"description": "",
	"eventTypes": [
	  "string"
	]
  }
`,
		Args: validators.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())
			var in []byte
			if len(args) > 0 {
				in = []byte(args[0])
			} else {
				var err error
				in, err = utils.ReadStdin()
				printer.CheckErr(err)
			}
			var ep wallet.CreateWebhookIn
			if len(in) > 0 {
				err := json.Unmarshal(in, &ep)
				printer.CheckErr(err)
			}

			// get flags
			if cmd.Flags().Changed(urlFlagName) {
				urlFlag, err := cmd.Flags().GetString(urlFlagName)
				printer.CheckErr(err)
				ep.Url = urlFlag
			}
			if cmd.Flags().Changed(eventTypesFlagName) {
				eventTypesFlag, err := cmd.Flags().GetStringArray(eventTypesFlagName)
				printer.CheckErr(err)
				ep.EventTypes = eventTypesFlag
			}
			if cmd.Flags().Changed(disabledFlagName) {
				disabledFlag, err := cmd.Flags().GetBool(disabledFlagName)
				printer.CheckErr(err)
				ep.Disabled = &disabledFlag
			}
			walletClient := getWalletClientOrExit()
			out, err := walletClient.Webhook.Create(cmd.Context(), &ep)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	create.Flags().String(urlFlagName, "", "")
	create.Flags().Int32(versionFlagName, 0, "")
	create.Flags().StringArray(eventTypesFlagName, []string{}, "")
	create.Flags().Bool(disabledFlagName, false, "")
	ec.cmd.AddCommand(create)

	// get
	get := &cobra.Command{
		Use:   "get ENDPOINT_ID",
		Short: "Get an endpoint by id",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			endpointID := args[0]

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Webhook.Retrieve(cmd.Context(), endpointID)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	ec.cmd.AddCommand(get)

	update := &cobra.Command{
		Use:   "update ENDPOINT_ID [JSON_PAYLOAD]",
		Short: "Update an application by id",
		Long: `Update an application by id

Example Schema:
{
  "url": "string",
  "version": 1,
  "description": "",
  "eventTypes": [
    "string"
  ]
}
`,
		Args: validators.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			// parse args
			endpointID := args[0]

			var in []byte
			if len(args) > 1 {
				in = []byte(args[1])
			} else {
				var err error
				in, err = utils.ReadStdin()
				printer.CheckErr(err)
			}
			var ep wallet.UpdateWebhookIn
			if len(in) > 0 {
				err := json.Unmarshal(in, &ep)
				printer.CheckErr(err)
			}

			// get flags
			if cmd.Flags().Changed(urlFlagName) {
				urlFlag, err := cmd.Flags().GetString(urlFlagName)
				printer.CheckErr(err)
				ep.Url = &urlFlag
			}
			if cmd.Flags().Changed(eventTypesFlagName) {
				eventTypesFlag, err := cmd.Flags().GetStringArray(eventTypesFlagName)
				printer.CheckErr(err)
				ep.EventTypes = eventTypesFlag
			}
			if cmd.Flags().Changed(disabledFlagName) {
				disabledFlag, err := cmd.Flags().GetBool(disabledFlagName)
				printer.CheckErr(err)
				ep.Disabled = &disabledFlag
			}

			walletClient := getWalletClientOrExit()
			out, err := walletClient.Webhook.Update(cmd.Context(), endpointID, &ep)
			printer.CheckErr(err)

			printer.Print(out)
		},
	}
	update.Flags().String(urlFlagName, "", "")
	update.Flags().Int32(versionFlagName, 1, "")
	update.Flags().StringArray(eventTypesFlagName, []string{}, "")
	update.Flags().Bool(disabledFlagName, false, "")
	ec.cmd.AddCommand(update)

	delete := &cobra.Command{
		Use:   "delete ENDPOINT_ID",
		Short: "Delete an endpoint by id",
		Args:  validators.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			// parse args
			endpointID := args[0]

			utils.Confirm(fmt.Sprintf("Are you sure you want to delete the the endpoint with id: %s", endpointID))

			walletClient := getWalletClientOrExit()
			err := walletClient.Webhook.Delete(cmd.Context(), endpointID)
			printer.CheckErr(err)

			fmt.Printf("Endpoint \"%s\" Deleted!\n", endpointID)
		},
	}
	ec.cmd.AddCommand(delete)

	/*
		secret := &cobra.Command{
			Use:   "secret ENDPOINT_ID",
			Short: "get an endpoint's secret by id",
			Args:  validators.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				printer := pretty.NewPrinter(getPrinterOptions())

				// parse args
				endpointID := args[0]

				walletClient := getWalletClientOrExit()
				out, err := walletClient.Webhook.GetSecret(cmd.Context(), endpointID)
				printer.CheckErr(err)

				printer.Print(out)
			},
		}
		ec.cmd.AddCommand(secret)
	*/

	return ec
}

func addEndpointFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("iterator", "i", "", "anchor id for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

func getEndpointListOptions(cmd *cobra.Command) *wallet.ListWebhookOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.ListWebhookOptions{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("iterator") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
