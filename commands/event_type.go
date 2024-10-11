package commands

import (
	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/spf13/cobra"
)

type eventTypeCmd struct {
	cmd *cobra.Command
}

func newEventTypeCmd() *eventTypeCmd {
	ec := &eventTypeCmd{}
	ec.cmd = &cobra.Command{
		Use:   "event-type",
		Short: "List, create & modify event-types",
	}

	// list
	list := &cobra.Command{
		Use:   "list",
		Short: "List current event-types",
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			apiClient := getAPIClientOrExit()
			l, err := apiClient.WebhookEventTypes.List(cmd.Context()) // , getEventTypeListOptions(cmd))
			printer.CheckErr(err)

			printer.Print(l)
		},
	}
	addEventTypeFilterFlags(list)
	ec.cmd.AddCommand(list)

	return ec
}

func addEventTypeFilterFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("iterator", "i", "", "anchor id for list call")
	cmd.Flags().Int32P("limit", "l", 50, "max items per request")
}

/*
func getEventTypeListOptions(cmd *cobra.Command) *wallet.ListWebhookOptions {
	limit, _ := cmd.Flags().GetInt32("limit")

	opts := &wallet.Li{
		Limit: &limit,
	}

	cursorFlag, _ := cmd.Flags().GetString("cursor")
	if cmd.Flags().Changed("iterator") {
		opts.Cursor = &cursorFlag
	}

	return opts
}
*/
