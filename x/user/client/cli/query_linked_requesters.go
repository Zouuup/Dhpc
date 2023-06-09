package cli

import (
	"context"

	"github.com/DhpcChain/Dhpc/x/user/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLinkedRequesters() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-linked-requesters",
		Short: "list all LinkedRequesters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLinkedRequestersRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LinkedRequestersAll(context.Background(), params)
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

func CmdShowLinkedRequesters() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-linked-requesters [index]",
		Short: "shows a LinkedRequesters",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetLinkedRequestersRequest{
				Index: argIndex,
			}

			res, err := queryClient.LinkedRequesters(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
