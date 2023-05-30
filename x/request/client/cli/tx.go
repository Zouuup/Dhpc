package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"Decent/x/request/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateAllowedOracles())
	cmd.AddCommand(CmdUpdateAllowedOracles())
	cmd.AddCommand(CmdDeleteAllowedOracles())
	cmd.AddCommand(CmdCreateMinerResponse())
	cmd.AddCommand(CmdUpdateMinerResponse())
	cmd.AddCommand(CmdDeleteMinerResponse())
	cmd.AddCommand(CmdCreateRequestRecord())
	cmd.AddCommand(CmdUpdateRequestRecord())
	cmd.AddCommand(CmdDeleteRequestRecord())
	cmd.AddCommand(CmdCreateAddTreasury())
	cmd.AddCommand(CmdUpdateAddTreasury())
	cmd.AddCommand(CmdDeleteAddTreasury())
	// this line is used by starport scaffolding # 1

	return cmd
}
