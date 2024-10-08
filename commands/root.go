package commands

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/openweb3-io/openwallet-cli/config"
	"github.com/openweb3-io/openwallet-cli/flags"
	"github.com/openweb3-io/openwallet-cli/version"
	wallet "github.com/openweb3-io/wallet-openapi/go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var defaultApiUrl = "https://api.wallet.openweb3.io"

const commandTimeout = 15 * time.Second

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "openwallet",
	Short:   "A CLI to interact with the Openweb3 Wallet API.",
	Version: version.Version,
}

func Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(version.String())

	// Root Flags
	rootCmd.Flags().BoolP("version", "v", false, "Get the version of the OpenWallet CLI") // overrides default msg

	// Global Flags
	color := "auto"
	colorFlag := flags.NewEnum(&color, "auto", "always", "never")
	flag.Var(colorFlag, "color", "auto|always|never")
	rootCmd.PersistentFlags().AddGoFlag(flag.Lookup("color"))
	cobra.CheckErr(viper.BindPFlag("color", rootCmd.PersistentFlags().Lookup("color"))) // allow color flag to be set in config

	// Register Commands
	rootCmd.AddCommand(newVersionCmd().cmd)
	rootCmd.AddCommand(newKeyCmd().cmd)
	rootCmd.AddCommand(newLoginCmd().cmd)
	rootCmd.AddCommand(newWalletCmd().cmd)
	// rootCmd.AddCommand(newAuthenticationCmd().cmd)
	// rootCmd.AddCommand(newOpenCmd().cmd)
	// rootCmd.AddCommand(newListenCmd().cmd)
	// rootCmd.AddCommand(newImportCmd().cmd)
	// rootCmd.AddCommand(newExportCmd().cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Setup config file
	configFolder, err := config.Folder()
	cobra.CheckErr(err)

	configFile := filepath.Join(configFolder, config.FileName)
	viper.SetConfigType("toml")
	viper.SetConfigFile(configFile)
	viper.SetConfigPermissions(config.FileMode)

	// read in environment variables that match
	viper.SetEnvPrefix("openwallet")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	_ = viper.ReadInConfig()
}

func getWalletClientOrExit() *wallet.WalletClient {
	apikey := viper.GetString("apikey")
	privateKey := viper.GetString("private_key")
	if privateKey == "" {
		fmt.Fprintln(os.Stderr, "No OPENWALLET_PRIVATE_KEY found!")
		fmt.Fprintln(os.Stderr, "Try running `openwallet login` to get started!")
		os.Exit(1)
	}

	opts := getWalletClientOptsOrExit(apikey, privateKey)
	return wallet.New(opts)
}

func getWalletClientOptsOrExit(apikey, privateKey string) *wallet.WalletClientOptions {
	opts := &wallet.WalletClientOptions{}
	rawServerUrl := viper.GetString("server_url")

	// fallback to debug_url for backwards compatibility
	if rawServerUrl == "" {
		rawServerUrl = viper.GetString("debug_url")
	}

	if rawServerUrl != "" {
		serverUrl, err := url.Parse(rawServerUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid server_url set: \"%s\"\n", rawServerUrl)
			os.Exit(1)
		}
		opts.ServerUrl = serverUrl
	}
	opts.ApiKey = apikey
	opts.PrivateKey = privateKey

	return opts
}
