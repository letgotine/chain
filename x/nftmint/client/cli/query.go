package cli

import (
	// "context"
	"context"
	"fmt"
	"strings"

	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryClassAttributes(),
		CmdQueryClassIdsByOwner(),
		CmdQueryClassIdsByName(),
		CmdQueryNFTMinter(),
	)

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), params)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassAttributes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-attributes",
		Args:  cobra.ExactArgs(1),
		Short: "Query the class attributes by class-id",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example: $ %s query %s class-attributes classID",
			version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassAttributes(
				context.Background(),
				&types.QueryClassAttributesRequest{ClassId: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassIdsByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-ids-by-owner",
		Args:  cobra.ExactArgs(1),
		Short: "Query classIDs owned by the owner address",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example: $ %s query %s class-ids-by-owner ununifi1exampleaddress",
			version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassIdsByOwner(
				context.Background(),
				&types.QueryClassIdsByOwnerRequest{Owner: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryClassIdsByName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "class-ids-by-name",
		Args:  cobra.ExactArgs(1),
		Short: "Query classIDs which have the class name",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Example: $ %s query %s class-name-id-list MyCollective",
			version.AppName, types.ModuleName),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClassIdsByName(
				context.Background(),
				&types.QueryClassIdsByNameRequest{ClassName: args[0]},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryNFTMinter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-minter",
		Args:  cobra.ExactArgs(2),
		Short: "Query nft minter with class and nft id",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.NFTMinter(
				context.Background(),
				&types.QueryNFTMinterRequest{
					ClassId: args[0],
					NftId:   args[1],
				},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}