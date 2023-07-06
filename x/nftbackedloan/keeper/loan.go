package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func (k Keeper) GetDebtByNft(ctx sdk.Context, nftIdBytes []byte) types.Loan {
	loan := types.Loan{}
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NftLoanKey(nftIdBytes))
	if bz == nil {
		return loan
	}

	k.cdc.MustUnmarshal(bz, &loan)
	return loan
}

func (k Keeper) GetAllDebts(ctx sdk.Context) []types.Loan {
	store := ctx.KVStore(k.storeKey)

	loans := []types.Loan{}
	it := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixNftLoan))
	defer it.Close()

	for ; it.Valid(); it.Next() {
		loan := types.Loan{}
		k.cdc.MustUnmarshal(it.Value(), &loan)

		loans = append(loans, loan)
	}
	return loans
}

func (k Keeper) SetDebt(ctx sdk.Context, loan types.Loan) {
	bz := k.cdc.MustMarshal(&loan)
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NftLoanKey(loan.NftId.IdBytes()), bz)
}

func (k Keeper) DeleteDebt(ctx sdk.Context, nftBytes []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NftLoanKey(nftBytes))
}

// remove debt (loan) from KVStore by using DeleteDebt method with the feature
// to judge if it exists before calling it
func (k Keeper) RemoveDebt(ctx sdk.Context, nftBytes []byte) {
	loan := k.GetDebtByNft(ctx, nftBytes)
	if !loan.Loan.Amount.IsNil() {
		k.DeleteDebt(ctx, nftBytes)
	}
}

func (k Keeper) IncreaseDebt(ctx sdk.Context, nftId types.NftIdentifier, amount sdk.Coin) {
	currDebt := k.GetDebtByNft(ctx, nftId.IdBytes())
	if sdk.Coin.IsNil(currDebt.Loan) {
		currDebt.NftId = nftId
		currDebt.Loan = amount
	} else {
		currDebt.Loan = currDebt.Loan.Add(amount)
	}
	k.SetDebt(ctx, currDebt)
}

func (k Keeper) DecreaseDebt(ctx sdk.Context, nftId types.NftIdentifier, amount sdk.Coin) {
	currDebt := k.GetDebtByNft(ctx, nftId.IdBytes())
	currDebt.Loan = currDebt.Loan.Sub(amount)
	k.SetDebt(ctx, currDebt)
}

func (k Keeper) Borrow(ctx sdk.Context, msg *types.MsgBorrow) error {
	bids := types.NftBids(k.GetBidsByNft(ctx, msg.NftId.IdBytes()))
	// if re-borrow, repay all borrowed bids
	borrowedBid := types.NftBids{}
	for _, bid := range bids {
		if len(bid.Borrowings) != 0 {
			borrowedBid = append(borrowedBid, bid)
		}
	}
	if len(borrowedBid) != 0 {
		err := k.AutoRepay(ctx, msg.NftId, borrowedBid, msg.Sender, msg.Sender)
		if err != nil {
			return err
		}
	}
	if !types.IsAbleToBorrow(bids, msg.BorrowBids, ctx.BlockTime()) {
		return types.ErrCannotBorrow
	}
	err := k.ManualBorrow(ctx, msg.NftId, msg.BorrowBids, msg.Sender, msg.Sender)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ManualBorrow(ctx sdk.Context, nft types.NftIdentifier, borrows []types.BorrowBid, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	borrowedAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, borrow := range borrows {
		if borrow.Amount.Denom != listing.BidDenom {
			return types.ErrInvalidBorrowDenom
		}
		bidderAddress, err := sdk.AccAddressFromBech32(borrow.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}
		usableAmount := bid.BorrowableAmount()
		if borrow.Amount.IsGTE(usableAmount) {
			borrowing := types.Borrowing{
				Amount:       usableAmount,
				LastRepaidAt: ctx.BlockTime(),
			}
			borrowedAmount = borrowedAmount.Add(usableAmount)
			bid.Borrowings = append(bid.Borrowings, borrowing)
		} else {
			borrowing := types.Borrowing{
				Amount:       borrow.Amount,
				LastRepaidAt: ctx.BlockTime(),
			}
			borrowedAmount = borrowedAmount.Add(borrow.Amount)
			bid.Borrowings = append(bid.Borrowings, borrowing)
		}
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	if !borrowedAmount.IsPositive() {
		return types.ErrInvalidBorrowAmount
	}

	k.IncreaseDebt(ctx, nft, borrowedAmount)

	receiverAddress, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddress, sdk.Coins{borrowedAmount})
	if err != nil {
		return err
	}

	// Emit event for borrow from bids
	_ = ctx.EventManager().EmitTypedEvent(&types.EventBorrow{
		Borrower: borrower,
		ClassId:  nft.ClassId,
		NftId:    nft.NftId,
		Amount:   borrowedAmount.String(),
	})

	return nil
}

func (k Keeper) Repay(ctx sdk.Context, msg *types.MsgRepay) error {
	return k.ManualRepay(ctx, msg.NftId, msg.RepayBids, msg.Sender, msg.Sender)
}

func (k Keeper) ManualRepay(ctx sdk.Context, nft types.NftIdentifier, repays []types.BorrowBid, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	sender, err := sdk.AccAddressFromBech32(borrower)
	if err != nil {
		return err
	}

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, listing.BidDenom)
	repayAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, repay := range repays {
		if repay.Amount.Denom != listing.BidDenom {
			return types.ErrInvalidRepayDenom
		}
		repayAmount = repayAmount.Add(repay.Amount)
	}
	if listerAmount.Amount.LT(repayAmount.Amount) {
		return types.ErrInsufficientBalance
	}

	currDebt := k.GetDebtByNft(ctx, nft.IdBytes())

	// return err if borrowing didn't happen once before
	if currDebt.Loan.IsNil() {
		return types.ErrNotBorrowed
	}

	repaidAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	repaidInterestAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, repay := range repays {
		bidderAddress, err := sdk.AccAddressFromBech32(repay.Bidder)
		if err != nil {
			return err
		}
		bid, err := k.GetBid(ctx, nft.IdBytes(), bidderAddress)
		if err != nil {
			return err
		}

		if len(bid.Borrowings) == 0 {
			continue
		}
		repaidAmount = repaidAmount.Add(repay.Amount)
		borrowings := []types.Borrowing{}
		for _, borrowing := range bid.Borrowings {
			if repay.Amount.IsZero() {
				break
			}
			receipt := borrowing.RepayThenGetReceipt(repay.Amount, ctx.BlockTime(), bid.CalcInterestF())
			repay.Amount = receipt.Charge
			// bid.InterestAmount = bid.InterestAmount.Add(receipt.PaidInterestAmount)
			repaidInterestAmount = repaidInterestAmount.Add(receipt.PaidInterestAmount)

			if !borrowing.IsAllRepaid() {
				borrowings = append(borrowings, borrowing)
			}
		}
		// clean up Borrowings
		bid.Borrowings = borrowings
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
		// Subtract when surplus repay amount exists
		repaidAmount = repaidAmount.Sub(repay.Amount)
	}

	repaidLoanAmount := repaidAmount.Sub(repaidInterestAmount)
	k.DecreaseDebt(ctx, nft, repaidLoanAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for repay
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: borrower,
		ClassId: nft.ClassId,
		NftId:   nft.NftId,
		Amount:  repaidAmount.String(),
	})
	return nil
}

func (k Keeper) AutoRepay(ctx sdk.Context, nft types.NftIdentifier, bids types.NftBids, borrower, receiver string) error {
	listing, err := k.GetNftListingByIdBytes(ctx, nft.IdBytes())
	if err != nil {
		return err
	}
	if listing.Owner != borrower {
		return types.ErrNotNftListingOwner
	}

	sender, err := sdk.AccAddressFromBech32(borrower)
	if err != nil {
		return err
	}

	listerAmount := k.bankKeeper.GetBalance(ctx, sender, listing.BidDenom)
	repayAmount, err := types.ExistRepayAmountAtTime(bids, ctx.BlockTime())
	if err != nil {
		return err
	}
	if listerAmount.Amount.LT(repayAmount.Amount) {
		return types.ErrInsufficientBalance
	}

	currDebt := k.GetDebtByNft(ctx, nft.IdBytes())

	// return err if borrowing didn't happen once before
	if currDebt.Loan.IsNil() {
		return types.ErrNotBorrowed
	}

	repaidAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	repaidInterestAmount := sdk.NewCoin(listing.BidDenom, sdk.ZeroInt())
	for _, bid := range bids {
		if len(bid.Borrowings) == 0 {
			continue
		}

		for _, borrowing := range bid.Borrowings {
			interest := bid.CalcInterest(borrowing.Amount, bid.InterestRate, borrowing.LastRepaidAt, ctx.BlockTime())
			repayAmount := borrowing.Amount.Add(interest)
			if repayAmount.IsZero() {
				break
			}
			repaidAmount = repaidAmount.Add(repayAmount)
			receipt := borrowing.RepayThenGetReceipt(repayAmount, ctx.BlockTime(), bid.CalcInterestF())
			// Subtract when surplus repay amount exists
			repaidAmount = repaidAmount.Sub(receipt.Charge)
			// bid.InterestAmount = bid.InterestAmount.Add(receipt.PaidInterestAmount)
			repaidInterestAmount = repaidInterestAmount.Add(receipt.PaidInterestAmount)
		}
		// clean up Borrowings
		bid.Borrowings = []types.Borrowing{}
		err = k.SetBid(ctx, bid)
		if err != nil {
			return err
		}
	}

	repaidLoanAmount := repaidAmount.Sub(repaidInterestAmount)
	k.DecreaseDebt(ctx, nft, repaidLoanAmount)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{repaidAmount})
	if err != nil {
		return err
	}

	// Emit event for repay
	_ = ctx.EventManager().EmitTypedEvent(&types.EventRepay{
		Repayer: borrower,
		ClassId: nft.ClassId,
		NftId:   nft.NftId,
		Amount:  repaidAmount.String(),
	})
	return nil
}
