// InMetrics represents the profit/loss amount in the metrics asset of the market.
// In the most cases, it means it's in "usd".
// And IMPORTANTLY, it means it's not calcualted in micro case.

package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) LiquidityProviderTokenRealAPY(c context.Context, req *types.QueryLiquidityProviderTokenRealAPYRequest) (*types.QueryLiquidityProviderTokenRealAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)

	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)
	return &types.QueryLiquidityProviderTokenRealAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) LiquidityProviderTokenNominalAPY(c context.Context, req *types.QueryLiquidityProviderTokenNominalAPYRequest) (*types.QueryLiquidityProviderTokenNominalAPYResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	rate := k.GetLPNominalYieldRate(ctx, req.BeforeHeight, req.AfterHeight)
	annualized := k.AnnualizeYieldRate(ctx, rate, req.BeforeHeight, req.AfterHeight)
	return &types.QueryLiquidityProviderTokenNominalAPYResponse{Apy: &annualized}, nil
}

func (k Keeper) PerpetualFutures(c context.Context, req *types.QueryPerpetualFuturesRequest) (*types.QueryPerpetualFuturesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	markets := k.GetParams(ctx).PerpetualFutures.Markets
	longPositions := sdk.ZeroDec()
	shortPositions := sdk.ZeroDec()
	for _, market := range markets {
		totalLongPositionSize := k.GetPerpetualFuturesPositionSizeInMetrics(ctx, *market, types.PositionType_LONG)

		if totalLongPositionSize.IsZero() {
			return nil, fmt.Errorf("long position size is zero")
		}
		totalShortPositionSize := k.GetPerpetualFuturesPositionSizeInMetrics(ctx, *market, types.PositionType_SHORT)
		if totalShortPositionSize.IsZero() {
			return nil, fmt.Errorf("short position size is zero")
		}
		longPositions.AddMut(totalLongPositionSize)
		shortPositions.AddMut(totalShortPositionSize)
	}

	metricsQuoteTicker := k.GetParams(ctx).PoolParams.QuoteTicker
	return &types.QueryPerpetualFuturesResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		LongPositions:      longPositions,
		ShortPositions:     shortPositions,
	}, nil
}

func (k Keeper) PerpetualFuturesMarket(c context.Context, req *types.QueryPerpetualFuturesMarketRequest) (*types.QueryPerpetualFuturesMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	price := sdk.NewDec(0)
	metricsQuoteTicker := k.GetParams(ctx).PoolParams.QuoteTicker
	market := types.Market{
		BaseDenom:  req.BaseDenom,
		QuoteDenom: req.QuoteDenom,
	}
	totalLongPositionSize := k.GetPerpetualFuturesPositionSizeInMetrics(ctx, market, types.PositionType_LONG)
	if totalLongPositionSize.IsZero() {
		return nil, fmt.Errorf("long position size is zero")
	}
	totalShortPositionSize := k.GetPerpetualFuturesPositionSizeInMetrics(ctx, market, types.PositionType_SHORT)
	if totalShortPositionSize.IsZero() {
		return nil, fmt.Errorf("short position size is zero")
	}

	return &types.QueryPerpetualFuturesMarketResponse{
		Price:              &price,
		MetricsQuoteTicker: metricsQuoteTicker,
		LongPositions:      &totalLongPositionSize,
		ShortPositions:     &totalShortPositionSize,
	}, nil
}

func (k Keeper) PerpetualOptions(c context.Context, req *types.QueryPerpetualOptionsRequest) (*types.QueryPerpetualOptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()

	return &types.QueryPerpetualOptionsResponse{}, nil
}

func (k Keeper) PerpetualOptionsMarket(c context.Context, req *types.QueryPerpetualOptionsMarketRequest) (*types.QueryPerpetualOptionsMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// TODO: implement the handler logic
	ctx.BlockHeight()

	return &types.QueryPerpetualOptionsMarketResponse{}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	metricsQuoteTicker := k.GetPoolQuoteTicker(ctx)
	poolMarketCap, err := k.GetPoolMarketCap(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryPoolResponse{
		MetricsQuoteTicker: metricsQuoteTicker,
		PoolMarketCap:      &poolMarketCap,
	}, nil
}

func (k Keeper) AllPositions(c context.Context, req *types.QueryAllPositionsRequest) (*types.QueryAllPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	var positions []*types.Position

	store := ctx.KVStore(k.storeKey)
	positionStore := prefix.NewStore(store, []byte(types.KeyPrefixPosition))
	pageRes, err := query.Paginate(positionStore, req.Pagination, func(key []byte, value []byte) error {
		position, err := k.UnmarshalPosition(value)
		if err != nil {
			return err
		}

		positions = append(positions, &position)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPositionsResponse{
		Positions:  positions,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) AddressPositions(c context.Context, req *types.QueryAddressPositionsRequest) (*types.QueryAddressPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Get all positions of NFTs owned by the address
	positions := k.GetAddressNFTPositions(ctx, address)

	queriedPositions, err := k.MakeQueriedPositions(ctx, positions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAddressPositionsResponse{
		Positions: queriedPositions,
	}, nil
}

func (k Keeper) MakeQueriedPositions(ctx sdk.Context, positions types.Positions) ([]types.QueriedPosition, error) {
	queriedPositions := make([]types.QueriedPosition, 0)
	usdMap := map[string]sdk.Dec{}
	for _, position := range positions {

		if _, ok := usdMap[position.Market.BaseDenom]; !ok {
			price, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
			if err != nil {
				return nil, err
			}
			usdMap[position.Market.BaseDenom] = price
		}

		if _, ok := usdMap[position.Market.QuoteDenom]; !ok {
			price, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
			if err != nil {
				return nil, err
			}
			usdMap[position.Market.QuoteDenom] = price
		}

		currentBaseUsdRate := usdMap[position.Market.BaseDenom]
		currentQuoteUsdRate := usdMap[position.Market.QuoteDenom]

		perpetualFuturesPosition, err := types.NewPerpetualFuturesPositionFromPosition(position)
		if err != nil {
			// not implemented perpetual options
			return nil, status.Error(codes.Internal, err.Error())
		}

		quoteTicker := k.GetPoolQuoteTicker(ctx)
		baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
		quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)
		profit := perpetualFuturesPosition.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate)
		if perpetualFuturesPosition.LeviedAmountNegative {
			profit = profit.Sub(perpetualFuturesPosition.LeviedAmountInMetrics(baseMetricsRate, quoteMetricsRate))
		} else {
			profit = profit.Add(perpetualFuturesPosition.LeviedAmountInMetrics(baseMetricsRate, quoteMetricsRate))
		}
		// fixme do not use sdk.Coin directly
		positiveOrNegativeProfitCoin := sdk.Coin{
			Denom:  "uusd",
			Amount: types.NormalToMicroInt(profit),
		}
		positiveOrNegativeEffectiveMargin := sdk.Coin{
			Denom:  "uusd",
			Amount: types.NormalToMicroInt(perpetualFuturesPosition.EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate)),
		}
		queriedPosition := types.QueriedPosition{
			Position:              position,
			ValuationProfit:       positiveOrNegativeProfitCoin,
			MarginMaintenanceRate: perpetualFuturesPosition.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate),
			EffectiveMargin:       positiveOrNegativeEffectiveMargin,
		}
		queriedPositions = append(queriedPositions, queriedPosition)
	}
	return queriedPositions, nil
}

func (k Keeper) Position(c context.Context, req *types.QueryPositionRequest) (*types.QueryPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	position := k.GetPositionWithId(ctx, req.PositionId)
	if position == nil {
		return &types.QueryPositionResponse{}, nil
	}

	perpetualFuturesPosition, err := types.NewPerpetualFuturesPositionFromPosition(*position)
	if err != nil {
		panic(err)
	}
	currentBaseUsdRate, err := k.GetCurrentPrice(ctx, position.Market.BaseDenom)
	if err != nil {
		panic(err)
	}
	currentQuoteUsdRate, err := k.GetCurrentPrice(ctx, position.Market.QuoteDenom)
	if err != nil {
		panic(err)
	}
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	baseMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.BaseDenom, currentBaseUsdRate)
	quoteMetricsRate := types.NewMetricsRateType(quoteTicker, position.Market.QuoteDenom, currentQuoteUsdRate)

	profit := perpetualFuturesPosition.ProfitAndLossInMetrics(baseMetricsRate, quoteMetricsRate)
	return &types.QueryPositionResponse{
		Position:              position,
		ValuationProfit:       sdk.NewCoin("uusd", types.NormalToMicroInt(profit)),
		MarginMaintenanceRate: perpetualFuturesPosition.MarginMaintenanceRate(baseMetricsRate, quoteMetricsRate),
		EffectiveMargin:       sdk.NewCoin("uusd", types.NormalToMicroInt(perpetualFuturesPosition.EffectiveMarginInMetrics(baseMetricsRate, quoteMetricsRate))),
	}, nil
}

func (k Keeper) AllPendingPaymentPositions(c context.Context, req *types.QueryAllPendingPaymentPositionsRequest) (*types.QueryAllPendingPaymentPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// Get all pending payment positions
	pendingPaymentPositions := k.GetAllPendingPaymentPositions(ctx)

	return &types.QueryAllPendingPaymentPositionsResponse{
		PendingPaymentPositions: pendingPaymentPositions,
	}, nil
}

func (k Keeper) PendingPaymentPosition(c context.Context, req *types.QueryPendingPaymentPositionRequest) (*types.QueryPendingPaymentPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	// Get all pending payment position by position id
	pendingPaymentPosition := k.GetPendingPaymentPosition(ctx, req.PositionId)

	return &types.QueryPendingPaymentPositionResponse{
		PendingPaymentPosition: pendingPaymentPosition,
	}, nil
}

func (k Keeper) PerpetualFuturesPositionSize(c context.Context, req *types.QueryPerpetualFuturesPositionSizeRequest) (*types.QueryPerpetualFuturesPositionSizeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Get all positions of NFTs owned by the address
	positions := types.Positions(k.GetAddressNFTPositions(ctx, address))
	getPriceFunc := func(ctx sdk.Context) func(denom string) (sdk.Dec, error) {
		return func(denom string) (sdk.Dec, error) {
			return k.GetCurrentPrice(ctx, denom)
		}
	}
	var result sdk.Dec
	quoteTicker := k.GetPoolQuoteTicker(ctx)
	if req.PositionType == types.PositionType_LONG {
		result, err = positions.EvaluateLongPositions(quoteTicker, getPriceFunc(ctx))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else if req.PositionType == types.PositionType_SHORT {
		result, err = positions.EvaluateShortPositions(quoteTicker, getPriceFunc(ctx))
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid position type")
	}
	return &types.QueryPerpetualFuturesPositionSizeResponse{
		TotalPositionSizeUsd: sdk.NewCoin("usd", result.RoundInt()),
	}, nil

}

func (k Keeper) DLPTokenRates(c context.Context, req *types.QueryDLPTokenRateRequest) (*types.QueryDLPTokenRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	var rates sdk.Coins
	for _, asset := range params.PoolParams.AcceptedAssetsConf {
		ldpDenomRate, err := k.LptDenomRate(ctx, asset.Denom)
		if err != nil {
			// todo error handing
			continue
		}
		// TODO: Is microzation necessary?
		// TODO: don't use NormalToMicroInt like this since it is hard to be consistent
		rates = append(rates, sdk.NewCoin(asset.Denom, types.NormalToMicroInt(ldpDenomRate)))
	}

	return &types.QueryDLPTokenRateResponse{
		Symbol: "DLP",
		Rates:  rates,
	}, nil
}

func (k Keeper) EstimateDLPTokenAmount(c context.Context, req *types.QueryEstimateDLPTokenAmountRequest) (*types.QueryEstimateDLPTokenAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	validAmount, ok := sdk.NewIntFromString(req.Amount)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	deposit := sdk.Coin{
		Denom:  req.MintDenom,
		Amount: validAmount,
	}
	if deposit.Validate() != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	depositFee, err := k.CalcDepositingFee(ctx, deposit, k.GetLPTokenBaseMintFee(ctx))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	depositFeeDeducted := deposit.Sub(depositFee)
	mintAmount, err := k.DetermineMintingLPTokenAmount(ctx, depositFeeDeducted)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryEstimateDLPTokenAmountResponse{
		EstimatedDlpAmount: mintAmount,
		DepositFee:         depositFee,
	}, nil
}

func (k Keeper) EstimateRedeemTokenAmount(c context.Context, req *types.QueryEstimateRedeemTokenAmountRequest) (*types.QueryEstimateRedeemTokenAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	validAmount, ok := sdk.NewIntFromString(req.LptAmount)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	validCoin := sdk.Coin{
		Denom:  req.RedeemDenom,
		Amount: validAmount,
	}
	if validCoin.Validate() != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	redeemAmount, redeemFee, err := k.GetRedeemDenomAmount(ctx, validCoin.Amount, validCoin.Denom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryEstimateRedeemTokenAmountResponse{
		Amount: redeemAmount,
		Fee:    redeemFee,
	}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) AvailableAssetInPoolByDenom(c context.Context, req *types.QueryAvailableAssetInPoolByDenomRequest) (*types.QueryAvailableAssetInPoolByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	availableAsset, err := k.AvailableAssetInPool(ctx, req.Denom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryAvailableAssetInPoolByDenomResponse{
		AvailableAsset: availableAsset,
	}, nil
}

func (k Keeper) AvailableAssetsInPool(c context.Context, req *types.QueryAvailableAssetsInPoolRequest) (*types.QueryAvailableAssetsInPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	availableAssets, err := k.AllAvailableAssetsInPool(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryAvailableAssetsInPoolResponse{
		AvailableAssets: availableAssets,
	}, nil
}
