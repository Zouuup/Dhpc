package cli

import (
	"strconv"

	"Decent/x/data/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-data [address] [owner] [network] [method] [score]",
		Short: "Broadcast message addData",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argOwner := args[1]
			argNetwork := args[2]
			argMethod := args[3]
			argScore, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddData(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argOwner,
				argNetwork,
				argMethod,
				argScore,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
