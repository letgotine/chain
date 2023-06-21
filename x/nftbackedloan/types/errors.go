package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNftListingAlreadyExists           = sdkerrors.Register(ModuleName, 1, "nft listing already exist")
	ErrNftListingDoesNotExist            = sdkerrors.Register(ModuleName, 2, "nft listing does not exist")
	ErrBidDoesNotExists                  = sdkerrors.Register(ModuleName, 3, "nft bid does not exist")
	ErrNotSupportedBidToken              = sdkerrors.Register(ModuleName, 4, "not supported bid token")
	ErrNftDoesNotExists                  = sdkerrors.Register(ModuleName, 5, "specified nft does not exist")
	ErrNotNftOwner                       = sdkerrors.Register(ModuleName, 6, "not the owner of nft")
	ErrNotNftListingOwner                = sdkerrors.Register(ModuleName, 7, "not the owner of nft listing")
	ErrNftBidAlreadyExists               = sdkerrors.Register(ModuleName, 8, "bid already exists on the nft")
	ErrNftBidDoesNotExists               = sdkerrors.Register(ModuleName, 9, "bid does not exists on the nft")
	ErrListingIsNotInStatusToBid         = sdkerrors.Register(ModuleName, 10, "listing is not in status to bid")
	ErrStatusCannotCancelListing         = sdkerrors.Register(ModuleName, 11, "listing is in the status where it can not be cancelled")
	ErrListingAlreadyEnded               = sdkerrors.Register(ModuleName, 12, "listing already ended")
	ErrNftListingNotInBidState           = sdkerrors.Register(ModuleName, 13, "listing is not on bid state")
	ErrInvalidBidDenom                   = sdkerrors.Register(ModuleName, 14, "invalid bid denom")
	ErrBidAlreadyExists                  = sdkerrors.Register(ModuleName, 15, "bid already exists")
	ErrNftListingNotInSuccessfulBidPhase = sdkerrors.Register(ModuleName, 16, "listing is not in successful bid status")
	ErrNftListingNotInLiquidation        = sdkerrors.Register(ModuleName, 17, "listing is not in liquidation status")
	ErrDebtExceedsMaxDebt                = sdkerrors.Register(ModuleName, 18, "debts exceeds maximum debt")
	ErrRepayAmountExceedsLoanAmount      = sdkerrors.Register(ModuleName, 19, "repay amount exceeds loan amount")
	ErrInvalidBorrowDenom                = sdkerrors.Register(ModuleName, 20, "invalid borrow denom")
	ErrInvalidRepayDenom                 = sdkerrors.Register(ModuleName, 21, "invalid repay denom")
	ErrNotTimeForCancel                  = sdkerrors.Register(ModuleName, 22, "not time for cancel yet")
	ErrInvalidBidAmount                  = sdkerrors.Register(ModuleName, 23, "invalid bid amount")
	ErrCannotCancelListingSingleBid      = sdkerrors.Register(ModuleName, 24, "cannot cancel single bid of the listing")
	ErrBidCancelIsAllowedAfterSomeTime   = sdkerrors.Register(ModuleName, 25, "bid cancel is allowed after some time after bid")
	ErrListingNeedsToBeBiddingStatus     = sdkerrors.Register(ModuleName, 26, "listing needs to be in BIDDING status")
	ErrNotBorrowed                       = sdkerrors.Register(ModuleName, 27, "not borrowed yet once")
	ErrBorrowedDeposit                   = sdkerrors.Register(ModuleName, 28, "borrowed bid deposit")
	ErrCannotCancelBid                   = sdkerrors.Register(ModuleName, 29, "listing is not in state where bid can be cancelled")
	ErrCannotParseDec                    = sdkerrors.Register(ModuleName, 30, "cannot parse decimal")
	ErrNotEnoughDeposit                  = sdkerrors.Register(ModuleName, 31, "not enough deposit")
	ErrBidParamInvalid                   = sdkerrors.Register(ModuleName, 32, "bid param invalid")
	ErrNotExistsBid                      = sdkerrors.Register(ModuleName, 33, "not exists bid")
	ErrCannotCancelListingWithBids       = sdkerrors.Register(ModuleName, 34, "cannot cancel listing with bids")
	ErrCannotCancelListingWithDebt       = sdkerrors.Register(ModuleName, 35, "cannot cancel listing with debt")
	ErrInterestAmountTooLarge            = sdkerrors.Register(ModuleName, 36, "interest amount too large")
	ErrInsufficientBalance               = sdkerrors.Register(ModuleName, 37, "insufficient balance")
	ErrSmallBiddingPeriod                = sdkerrors.Register(ModuleName, 38, "bidding period is too short")
	ErrNotExistsNft                      = sdkerrors.Register(ModuleName, 39, "not exists nft")
	ErrMinimumDepositRateTooHigh         = sdkerrors.Register(ModuleName, 40, "minimum deposit rate too high")
	ErrCannotBorrow                      = sdkerrors.Register(ModuleName, 41, "cannot borrow from bids")
	ErrAlreadyBorrowed                   = sdkerrors.Register(ModuleName, 42, "already borrowed")
)
