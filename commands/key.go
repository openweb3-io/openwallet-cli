package commands

import (
	"crypto/ed25519"
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
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "Failed to generate private key, please try again.")
		os.Exit(1)

	}

	publicKeyStr := hex.EncodeToString(publicKey)
	privateKeyStr := hex.EncodeToString(privateKey)

	fmt.Println("Key generated")
	fmt.Printf("Private Key(hex): %s\n", privateKeyStr)
	fmt.Printf("Public Key(hex): %s\n", publicKeyStr)
}
