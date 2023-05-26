package cli

import (
	"Decent/x/request/types"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateRequestRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-request-record [uuid] [t-xhash] [network] [from] [address] [score] [oracle] [block] [stage] [miners]",
		Short: "Create a new requestRecord",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUUID := args[0]

			// Get value arguments
			argTXhash := args[1]
			argNetwork := args[2]
			argFrom := args[3]
			argAddress := args[4]
			argScore, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}
			argOracle := args[6]
			argBlock, err := cast.ToInt32E(args[7])
			if err != nil {
				return err
			}
			argStage, err := cast.ToInt32E(args[8])
			if err != nil {
				return err
			}
			argMiners := new(types.MinerResponse)
			err = json.Unmarshal([]byte(args[9]), argMiners)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRequestRecord(
				clientCtx.GetFromAddress().String(),
				indexUUID,
				argTXhash,
				argNetwork,
				argFrom,
				argAddress,
				argScore,
				argOracle,
				argBlock,
				argStage,
				argMiners,
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

func CmdUpdateRequestRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-request-record [uuid] [t-xhash] [network] [from] [address] [score] [oracle] [block] [stage] [miners]",
		Short: "Update a requestRecord",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexUUID := args[0]

			// Get value arguments
			argTXhash := args[1]
			argNetwork := args[2]
			argFrom := args[3]
			argAddress := args[4]
			argScore, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}
			argOracle := args[6]
			argBlock, err := cast.ToInt32E(args[7])
			if err != nil {
				return err
			}
			argStage, err := cast.ToInt32E(args[8])
			if err != nil {
				return err
			}
			argMiners := new(types.MinerResponse)
			err = json.Unmarshal([]byte(args[9]), argMiners)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRequestRecord(
				clientCtx.GetFromAddress().String(),
				indexUUID,
				argTXhash,
				argNetwork,
				argFrom,
				argAddress,
				argScore,
				argOracle,
				argBlock,
				argStage,
				argMiners,
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

func CmdDeleteRequestRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-request-record [uuid]",
		Short: "Delete a requestRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexUUID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRequestRecord(
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
