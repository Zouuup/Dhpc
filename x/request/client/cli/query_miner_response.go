package cli

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/request/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListMinerResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-miner-response",
		Short: "list all minerResponse",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMinerResponseRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MinerResponseAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowMinerResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-miner-response [uuid]",
		Short: "shows a minerResponse",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUUID := args[0]

			params := &types.QueryGetMinerResponseRequest{
				UUID: argUUID,
			}

			res, err := queryClient.MinerResponse(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
