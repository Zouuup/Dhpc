package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/DhpcChain/Dhpc/x/request/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group request queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListAllowedOracles())
	cmd.AddCommand(CmdShowAllowedOracles())
	cmd.AddCommand(CmdListMinerResponse())
	cmd.AddCommand(CmdShowMinerResponse())
	cmd.AddCommand(CmdListRequestRecord())
	cmd.AddCommand(CmdListMinerPendingRequestRecord())
	cmd.AddCommand(CmdListAnswerPendingRequestRecord())
	cmd.AddCommand(CmdShowRequestRecord())
	cmd.AddCommand(CmdShowTreasury())
	// this line is used by starport scaffolding # 1

	return cmd
}
