package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	codecType "github.com/cosmos/cosmos-sdk/codec/types"

	ununifiType "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		CmdMintLiquidityProviderToken(),
		CmdBurnLiquidityProviderToken(),
		CmdOpenPosition(),
		CmdClosePosition(),
		CmdReportLiquidation(),
	)

	return cmd
}

func CmdMintLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-lpt [amount]",
		Short: "mint liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`mint liquidity provider token.
Example:
$ %s tx %s mint-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.MsgMintLiquidityProviderToken{
				Sender: ununifiType.StringAccAddress(sender),
				Amount: amount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdBurnLiquidityProviderToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-lpt [amount]",
		Short: "burn liquidity provider token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`burn liquidity provider token.
Example:
$ %s tx %s burn-lpt --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()
			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid amount")
			}

			msg := types.MsgBurnLiquidityProviderToken{
				Sender: ununifiType.StringAccAddress(sender),
				Amount: amount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdOpenPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "open-position",
		Short:                      fmt.Sprintf("%s open position subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(
		CmdOpenPerpetualFuturesPosition(),
		CmdOpenPerpetualOptionsPosition(),
	)

	return cmd
}

func CmdOpenPerpetualFuturesPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-futures [margin] [base-denom] [quote-denom]",
		Short: "open perpetual futures position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual futures position.
Example:
$ %s tx %s open-position perpetual-futures --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()
			margin, err := sdk.ParseCoinNormalized(args[0])
			baseDenom := args[1]
			quoteDenom := args[2]
			if err != nil {
				return err
			}

			// todo: use args to create position instance
			positionInstVal := types.PerpetualFuturesPositionInstance{
				PositionType: types.PositionType_LONG,
				Size_:        sdk.NewDecWithPrec(100, 0),
				Leverage:     5,
			}
			positionInstBz, err := positionInstVal.Marshal()
			if err != nil {
				return err
			}
			positionInstance := codecType.Any{
				TypeUrl: "/ununifi.derivatives.PerpetualFuturesPositionInstance",
				Value:   positionInstBz,
			}

			market := types.Market{
				BaseDenom:  baseDenom,
				QuoteDenom: quoteDenom,
			}
			msg := types.NewMsgOpenPosition(sender, margin, market, positionInstance)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdOpenPerpetualOptionsPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "perpetual-options [margin] [base-denom] [quote-denom]",
		Short: "open perpetual options position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`open perpetual options position.
Example:
$ %s tx %s open-position perpetual-options --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented")
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdClosePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`close position.
Example:
$ %s tx %s close-position --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgClosePosition{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdReportLiquidation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-liquidation",
		Short: "report liquidation needed position",
		Long: strings.TrimSpace(
			fmt.Sprintf(`report liquidation needed position.
Example:
$ %s tx %s report-liquidation --from myKeyName --chain-id ununifi-x
`, version.AppName, types.ModuleName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress()

			msg := types.MsgReportLiquidation{
				Sender: ununifiType.StringAccAddress(sender),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
