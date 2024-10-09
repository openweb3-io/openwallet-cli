package commands

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/openweb3-io/openwallet-cli/validators"
	"github.com/spf13/cobra"
)

type keyCmd struct {
	cmd *cobra.Command
}

func newKeyCmd() *keyCmd {
	lc := &keyCmd{}
	lc.cmd = &cobra.Command{
		Use:   "key",
		Short: "Generete a OpenWallet Private Key",
		Args:  validators.NoArgs(),
	}

	// generate
	generate := &cobra.Command{
		Use:   "generate",
		Short: "Generate private Key",
		Run:   lc.generate,
	}

	lc.cmd.AddCommand(generate)

	return lc
}

func (lc *keyCmd) generate(cmd *cobra.Command, args []string) {
	apikey, secret, err := GenerateApiKey()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "Failed to generate private key, please try again.")
		os.Exit(1)

	}

	fmt.Println("Key generated")
	fmt.Printf("apikey: %s\n", apikey)
	fmt.Printf("secret: %s\n", secret)
}

func GenerateApiKey() (apiKey, apiSecret string, err error) {
	pk, sk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	apiSecret = hex.EncodeToString(sk.Seed())
	apiKey = hex.EncodeToString(pk)
	return
}
