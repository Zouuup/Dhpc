package cli

import (
	"context"
	"strconv"

	"github.com/DhpcChain/Dhpc/x/request/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListAllowedOracles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-allowed-oracles",
		Short: "list all allowedOracles",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAllowedOraclesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllowedOraclesAll(context.Background(), params)
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

func CmdShowAllowedOracles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-allowed-oracles [id]",
		Short: "shows a allowedOracles",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAllowedOraclesRequest{
				Id: id,
			}

			res, err := queryClient.AllowedOracles(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
