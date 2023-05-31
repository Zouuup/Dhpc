package cli

import (
	"context"

	"Dhpc/x/request/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListRequestRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-request-record",
		Short: "list all requestRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRequestRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RequestRecordAll(context.Background(), params)
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

func CmdShowRequestRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-request-record [uuid]",
		Short: "shows a requestRecord",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUUID := args[0]

			params := &types.QueryGetRequestRecordRequest{
				UUID: argUUID,
			}

			res, err := queryClient.RequestRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
