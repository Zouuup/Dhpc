package cli

import (
	"Decent/x/request/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateMinerResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-miner-response [uuid] [request-uuid] [hash] [answer] [dataUsed]",
		Short: "Create a new minerResponse",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUUID := args[0]

			// Get value arguments
			argRequestUUID := args[1]
			argHash := args[2]
			argAnswer, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argDataUsed := args[4]

			msg := types.NewMsgCreateMinerResponse(
				clientCtx.GetFromAddress().String(),
				indexUUID,
				argRequestUUID,
				argHash,
				argAnswer,
				argDataUsed,
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

func CmdUpdateMinerResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-miner-response [uuid] [request-uuid] [hash] [answer] [salt]",
		Short: "Update a minerResponse",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUUID := args[0]

			// Get value arguments
			argRequestUUID := args[1]
			argHash := args[2]
			argAnswer, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argDataUsed := args[4]
			argSalt, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMinerResponse(
				clientCtx.GetFromAddress().String(),
				indexUUID,
				argRequestUUID,
				argHash,
				argAnswer,
				argDataUsed,
				argSalt,
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

func CmdDeleteMinerResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-miner-response [uuid]",
		Short: "Delete a minerResponse",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexUUID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMinerResponse(
				clientCtx.GetFromAddress().String(),
				indexUUID,
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
