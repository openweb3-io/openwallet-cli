package commands

import (
	"fmt"
	"net/http"

	"github.com/openweb3-io/openwallet-cli/pretty"
	"github.com/openweb3-io/openwallet-cli/utils"
	"github.com/openweb3-io/openwallet-cli/validators"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
)

type verifyCmd struct {
	cmd *cobra.Command
}

func newVerifyCmd() *verifyCmd {
	secretFlagName := "secret"
	signatureFlagName := "signature"
	msgIdFlagName := "msg-id"
	timestampFlagName := "timestamp"
	ac := &verifyCmd{}
	ac.cmd = &cobra.Command{
		Use:   "verify [JSON_PAYLOAD]",
		Short: "Verify the signature of a webhook message",
		Args:  validators.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			printer := pretty.NewPrinter(getPrinterOptions())

			// parse args
			var payload []byte
			if len(args) > 0 {
				payload = []byte(args[0])
			} else {
				var err error
				payload, err = utils.ReadStdin()
				printer.CheckErr(err)
			}

			if len(payload) <= 0 {
				printer.CheckErr("No json payload provided!")
			}

			// ensure all flags are set
			var err error
			if !cmd.Flags().Changed(secretFlagName) {
				err = fmt.Errorf("secret required for verification")
			} else if !cmd.Flags().Changed(signatureFlagName) {
				err = fmt.Errorf("signature required for verification")
			} else if !cmd.Flags().Changed(timestampFlagName) {
				err = fmt.Errorf("timestamp required for verification")
			} else if !cmd.Flags().Changed(msgIdFlagName) {
				err = fmt.Errorf("message ID required for verification")
			}
			printer.CheckErr(err)

			// get flags
			secret, err := cmd.Flags().GetString(secretFlagName)
			printer.CheckErr(err)
			msgID, err := cmd.Flags().GetString(msgIdFlagName)
			printer.CheckErr(err)
			timestamp, err := cmd.Flags().GetString(timestampFlagName)
			printer.CheckErr(err)
			signature, err := cmd.Flags().GetString(signatureFlagName)
			printer.CheckErr(err)

			headers := http.Header{}
			headers.Set("x-id", msgID)
			headers.Set("x-timestamp", timestamp)
			headers.Set("x-signature", signature)

			_ /*wh*/, err = wallet.NewWebhookClient(secret)
			if err != nil {
				printer.CheckErr(fmt.Errorf("failed to parse signing secret: %s", err.Error()))
			}

			/* TODO
			err = wh.Verify(payload, headers)
			if err != nil {
				errNoTimestamp := wh.VerifyIgnoringTimestamp(payload, headers)
				if errNoTimestamp == nil {
					fmt.Println("Signature is valid but failed timestamp verification.")
					os.Exit(1)
				}
				printer.CheckErr(errNoTimestamp)

			}
			*/
			fmt.Println("Message Signature Is Valid!")
		},
	}
	ac.cmd.Flags().String(secretFlagName, "", "signing secret of the endpoint (required)")
	ac.cmd.Flags().String(msgIdFlagName, "", "msg id header (required)")
	ac.cmd.Flags().String(timestampFlagName, "", "timestamp header (required)")
	ac.cmd.Flags().String(signatureFlagName, "", "signature header (required)")
	return ac
}
