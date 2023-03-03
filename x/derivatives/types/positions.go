package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type PositionInstance interface {
	proto.Message
}

func UnpackPositionInstance(positionAny types.Any) (PositionInstance, error) {
	position := UnpackPerpetualFuturesPositionInstance(positionAny)
	if position != nil {
		return position, nil
	}

	position = UnpackPerpetualOptionsPosition(positionAny)
	if position != nil {
		return position, nil
	}

	return nil, fmt.Errorf("this Any doesn't have PositionInstance value")
}

func MustUnpackPositionInstance(positionAny types.Any) PositionInstance {
	position, err := UnpackPositionInstance(positionAny)
	if err != nil {
		panic(err)
	}
	return position
}

func (m Position) NeedLiquidation(MarginMaintenanceRate, baseClosedCurrency, quoteCurrentRate sdk.Dec) bool {
	ins, err := UnpackPositionInstance(m.PositionInstance)
	if err != nil {
		return false
	}

	switch positionInstance := ins.(type) {
	case *PerpetualFuturesPositionInstance:
		perpetualFuturesPosition := NewPerpetualFuturesPosition(m, *positionInstance)
		return perpetualFuturesPosition.NeedLiquidation(MarginMaintenanceRate, baseClosedCurrency, quoteCurrentRate)
		break
	case *PerpetualOptionsPositionInstance:
		panic("not implemented")
		break
	default:
		panic("not implemented")
	}
	return false
}

func NewPerpetualFuturesPosition(position Position, ins PerpetualFuturesPositionInstance) PerpetualFuturesPosition {
	return PerpetualFuturesPosition{
		Id:               position.Id,
		Market:           position.Market,
		Address:          position.Address,
		OpenedAt:         position.OpenedAt,
		OpenedBaseRate:   position.OpenedBaseRate,
		OpenedQuoteRate:  position.OpenedQuoteRate,
		OpenedHeight:     position.OpenedHeight,
		RemainingMargin:  position.RemainingMargin,
		LastLeviedAt:     position.LastLeviedAt,
		PositionInstance: ins,
	}
}

func (m PerpetualFuturesPosition) NeedLiquidation(minMarginMaintenanceRate, baseClosedCurrency, quoteCurrentRate sdk.Dec) bool {
	marginMaintenanceRate := m.GetMarginMaintenanceRate(baseClosedCurrency, quoteCurrentRate)
	if marginMaintenanceRate.LT(minMarginMaintenanceRate) {
		return true
	} else {
		return false
	}
}

func (m PerpetualFuturesPosition) GetMarginMaintenanceRate(baseCurrentRate, quoteCurrentRate sdk.Dec) sdk.Dec {
	if m.PositionInstance.PositionType == PositionType_LONG {
		marginRequirement := m.PositionInstance.MarginRequirement(m.OpenedBaseRate)
		fmt.Println(marginRequirement.String())
		effectiveMargin := sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(baseCurrentRate)
		fmt.Println(effectiveMargin.String())
		marginMaintenanceRate := effectiveMargin.Quo(marginRequirement)
		fmt.Println(marginMaintenanceRate.String())
		return marginMaintenanceRate
	} else {
		// case position type is short
		marginRequirement := m.PositionInstance.MarginRequirement(quoteCurrentRate)
		effectiveMargin := sdk.NewDecFromInt(m.RemainingMargin.Amount).Mul(m.OpenedQuoteRate)
		marginMaintenanceRate := effectiveMargin.Quo(marginRequirement)
		return marginMaintenanceRate
	}
}

func (m PerpetualFuturesPosition) CalcProfit(closedRate sdk.Dec) Revenue {
	sub := closedRate.Sub(m.OpenedBaseRate)
	revenue := m.GetRevenueType(sub)
	if sub.IsNegative() {
		sub = sub.Neg()
	}
	// resultDec := sub.Mul(m.PositionInstance.GetOrderSize())

	resultDec := sub.Mul(m.PositionInstance.Size_)

	// profit must be calculated in remaining margin denom
	if m.RemainingMargin.Denom != m.Market.QuoteDenom {
		resultDec = resultDec.Quo(closedRate)
	}

	// make it micro unit by multiplying 1000000
	// this means it assumes the price difference is calculated in normal unit, not micro unit.
	// e.g. In ubtc/uusdc market, the market price of ubtc is actually in BTC unit.
	// And, the position size follows the market price unit.
	actualResultAmount := resultDec.Mul(sdk.MustNewDecFromStr("1000000")).RoundInt()

	return Revenue{
		RevenueType: revenue,
		Amount:      sdk.NewCoin(m.RemainingMargin.Denom, actualResultAmount),
	}
}

func (m PerpetualFuturesPosition) GetRevenueType(sub sdk.Dec) RevenueType {
	if m.PositionInstance.PositionType == PositionType_LONG {
		if sub.IsPositive() {
			return RevenueType_PROFIT
		} else {
			return RevenueType_LOSS
		}
	} else if m.PositionInstance.PositionType == PositionType_SHORT {
		// todo: think about amount is zero case
		if sub.IsNegative() {
			return RevenueType_PROFIT
		} else {
			return RevenueType_LOSS
		}
	} else {
		panic("not implemented")
	}
}

func (a Revenue) Equal(b Revenue) bool {
	if a.RevenueType != b.RevenueType {
		return false
	}
	if !a.Amount.IsEqual(b.Amount) {
		return false
	}
	return true
}
