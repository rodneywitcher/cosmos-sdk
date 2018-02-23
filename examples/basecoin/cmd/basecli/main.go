package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/tendermint/tmlibs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/version"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/commands"
)

// gaiacliCmd is the entry point for this binary
var (
	basecliCmd = &cobra.Command{
		Use:   "basecli",
		Short: "Basecoin light-client",
	}

	lineBreak = &cobra.Command{Run: func(*cobra.Command, []string) {}}
)

func todoNotImplemented(_ *cobra.Command, _ []string) error {
	return errors.New("TODO: Command not yet implemented")
}

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// generic client commands
	AddClientCommands(basecliCmd)

	// query/post commands (custom to binary)
	basecliCmd.AddCommand(
		client.GetCommands(
			bankcmd.GetAccountCmd("main"),
		)...)
	basecliCmd.AddCommand(
		client.PostCommands(
			bankcmd.SendTxCommand(),
		)...)

	// add proxy, version and key info
	basecliCmd.AddCommand(
		lineBreak,
		serveCommand(),
		keys.Commands(),
		lineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(basecliCmd, "BC", os.ExpandEnv("$HOME/.basecli"))
	executor.Execute()
}