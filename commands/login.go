package commands

import (
	"fmt"
	"net/url"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/openweb3-io/openwallet-cli/config"
	"github.com/openweb3-io/openwallet-cli/validators"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type loginCmd struct {
	cmd *cobra.Command
}

func newLoginCmd() *loginCmd {
	lc := &loginCmd{}
	lc.cmd = &cobra.Command{
		Use:   "login",
		Short: "Interactively configure your OpenWallet API credentials",
		Args:  validators.NoArgs(),
		Run:   lc.run,
	}
	return lc
}

func (lc *loginCmd) run(cmd *cobra.Command, args []string) {
	fmt.Printf("Welcome to the OpenWallet CLI, enter your auth token to get started!\n\n")

	defaultServerUrl := viper.GetString("server_url")
	if defaultServerUrl == "" {
		defaultServerUrl = defaultApiUrl
	}

	// get server_url
	serverUrlPrompt := promptui.Prompt{
		Label:   "OpenWallet Server URL",
		Default: defaultServerUrl,
	}
	serverUrl, err := serverUrlPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initialization failed %v\n", err)
		os.Exit(1)
	}
	if _, err := url.Parse(serverUrl); err != nil {
		fmt.Fprintf(os.Stderr, "Invalid server url %s\n%v\n", serverUrl, err)
		os.Exit(1)
	}
	if serverUrl != defaultServerUrl && serverUrl != "" {
		viper.Set("server_url", serverUrl)
	}

	// get apikey
	defaultApikey := viper.GetString("apikey")
	apikeyPrompt := promptui.Prompt{
		Label:   "OpenWallet API Key",
		Default: defaultApikey,
	}
	apikey, err := apikeyPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initialization failed %v\n", err)
		os.Exit(1)
	}
	if apikey != defaultApikey && apikey != "" {
		viper.Set("apikey", apikey)
	}

	// get private key
	defaultPrivateKey := viper.GetString("private_key")
	privateKeyPrompt := promptui.Prompt{
		Label:   "OpenWallet Private Key",
		Default: defaultPrivateKey,
	}
	privateKey, err := privateKeyPrompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initialization failed %v\n", err)
		os.Exit(1)
	}
	if privateKey != defaultPrivateKey && privateKey != "" {
		viper.Set("private_key", privateKey)
	}

	if err := config.Write(viper.AllSettings()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		fmt.Fprintln(os.Stderr, "Failed to configure the OpenWallet CLI, please try again or try setting your auth token manually 'OPENWALLET_PRIVATE_KEY' environment variable.")
		os.Exit(1)
	}

	fmt.Printf("All Set! Your config has been written to \"%s\"\n", viper.ConfigFileUsed())
	fmt.Println("Type `openwallet --help` to print the OpenWallet CLI documentation!")
}
