package cli

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/data/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-data",
		Short: "list all data",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DataAll(context.Background(), params)
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

func CmdListDataByAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-data-by-addr [address]",
		Short: "list all data by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			address := args[0]
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDataRequestByAddr{
				Address:    address,
				Pagination: pageReq,
			}

			res, err := queryClient.DataAllByAddr(context.Background(), params)
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

func CmdShowData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-data [address] [owner] [network]",
		Short: "shows a data",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argOwner := args[1]
			argNetwork := args[2]

			params := &types.QueryGetDataRequest{
				Address: argAddress,
				Owner:   argOwner,
				Network: argNetwork,
			}

			res, err := queryClient.Data(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
