<h1 align="center">
  <a href="https://wallet.openweb3.io">
    <p align="center">OpenWallet - Wallet as a service</p>
  </a>
</h1>

# OpenWallet CLI

[![GitHub release (latest by date)][release-img]][release]
[![GolangCI][golangci-lint-img]][golangci-lint]
[![Go Report Card][report-card-img]][report-card]

A CLI to interact with the OpenWallet API.

**With the OpenWallet CLI, you can:**

- Interact with the OpenWallet CLI

## Installation

### macOS

The OpenWallet CLI is available on macOS via [Homebrew](https://brew.sh/):

```sh
brew install openweb3-io/openwallet
```

### Windows

The OpenWallet CLI is available on Windows via the [Scoop](https://scoop.sh/) package manager:

```sh
scoop bucket add openwallet https://github.com/openweb3-io/scoop-openwallet.git
scoop install openwallet
```

### Linux

The OpenWallet CLI is available on Linux via:

* The [Snap Store](https://snapcraft.io): `snap install openwallet`
* The [Arch User Repository (AUR)](https://wiki.archlinux.org/title/Arch_User_Repository): `yay -S openwallet-cli`
* For Ubuntu/Debian: get the `deb` package from [our Github releases page](https://github.com/openweb3-io/openwallet-cli/releases)
* For Fedora/CentOS: get the `rpm` package from [our Github releases page](https://github.com/openweb3/openwallet-cli/releases)


### Pre-built executables

You can download and use our pre-built executables directly from [our releases page](https://github.com/openweb3-io/openwallet-cli/releases), and use them as is without having to install anything.

1. Download and extract the `tar.gz` archive for your operating system.
2. Run the `openwallet` executable from the command line: `./openwallet help`.

Note: you may need to allow execution by running `chmod +x openwallet`.


You can also put the binaries anywhere in your `PATH` so you can run the command from anywhere without needing to provide its full path. On macOS or Linux you can achieve this by moving the executable to `/usr/local/bin` or `/usr/bin`.


## Usage

Installing the OpenWallet CLI provides access to the `openwallet` command.

```sh
openwallet [command]

# Run `openwallet help` for information about the available commands
openwallet help

# or add the `--help` flag to any command for a more detailed description and list of flags
openwallet [command] --help
```


## Using the `listen` command

The `listen` command creates an on-the-fly publicly accessible URL for use when testing webhooks.

**NOTE:** You don't need a OpenWallet account when using the `listen` command.

The cli then acts as a proxy, forwarding any requests to the given local URL.
This is useful for testing your webhook server locally without having to open a port or
change any NAT configuration on your network.

Example:

`openwallet listen http://localhost:8000/webhook/`

## Interacting with the OpenWallet server

```sh
# Set your Secret temporarily via the OPENWALLET_SECRET environment variable
export OPENWALLET_SECRET=<MY-SECRET>
# or to persistently store your auth token in a config file run
openwallet login # interactively configure your OpenWallet API credentials

# Create an Wallet with the name "Demo"
openwallet wallet create '{ "name": "demo" }'
# or pipe in some json
echo '{ "name": "demo" }' | openwallet wallet create
# or use the convenience cli flags
openwallet wallet create --data-name demo

# List Wallets
openwallet wallet list --limit 2 --cursor some_cursor 
```

## Commands

The OpenWallet CLI supports the following commands:
| Command         | Description                                                |
| --------------- | ---------------------------------------------------------- |
| login           | Interactively configure your OpenWallet API credentials    |
| wallet          | List, create & modify wallets                              |
| authentication  | Manage authentication tasks such as getting dashboard URLs |
| endpoint        | List, create & modify endpoints                            |
| event-type      | List, create & modify event types                          |
| verify          | Verify the signature of a webhook message                  |
| listen          | Forward webhook requests a local url                       |
| integration     | List, create & modify integrations                         |
| import          | Import data from a file to your OpenWallet Organization    |
| export          | Export data from your OpenWallet Organization to a file    |
| open            | Quickly open OpenWallet pages in your browser              |
| completion      | Generate completion script                                 |
| version         | Get the version of the OpenWallet CLI                      |
| help            | Help about any command                                     |


## Shell Completions

Shell completion scripts are provided for Bash, Zsh, fish, & PowerShell.

To generate a script for your shell type `openwallet completion <SHELL NAME>`.

For detailed instructions on configuring completions for your shell run `openwallet completion --help`.


## Documentation

For a more information, checkout our [API reference](https://docs.openweb3.io).


### Development

#### Building the current commit

This project uses [goreleaser](https://github.com/goreleaser/goreleaser/).

 1) Install [go](https://golang.org/doc/install).
 2) Install [snapcraft](https://snapcraft.io/docs/installing-snapcraft).
 3) Install goreleaser via the steps [here](https://goreleaser.com/install/).
 4) Build current commit via `goreleaser release --snapshot --skip-publish --rm-dist`.

[release-img]: https://img.shields.io/github/v/release/openweb3-io/openwallet-cli
[release]: https://github.com/openweb3-io/openwallet-cli/releases
[golangci-lint-img]: https://github.com/openweb3-io/openwallet-cli/workflows/go-lint/badge.svg
[golangci-lint]: https://github.com/openweb3-io/openwallet-cli/actions?query=workflow%3Ago-lint
[report-card-img]: https://goreportcard.com/badge/github.com/openweb3-io/openwallet-cli
[report-card]: https://goreportcard.com/report/github.com/openweb3-io/openwallet-cli