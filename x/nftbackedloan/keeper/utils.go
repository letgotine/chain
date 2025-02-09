package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftbackedloan/types"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func validateListNftMsg(k Keeper, ctx sdk.Context, msg *types.MsgListNft) error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return checkListNft(k, ctx, sender, msg.NftId, msg.BidDenom, msg.MinimumDepositRate)
}

func checkListNft(k Keeper, ctx sdk.Context, sender sdk.AccAddress, nftId types.NftIdentifier, bidToken string, minimumDepositRate sdk.Dec) error {
	// check listing already exists
	_, err := k.GetNftListingByIdBytes(ctx, nftId.IdBytes())
	if err == nil {
		return types.ErrNftListingAlreadyExists
	}

	// Check nft exists
	_, found := k.nftKeeper.GetNFT(ctx, nftId.ClassId, nftId.NftId)
	if !found {
		return types.ErrNftDoesNotExists
	}

	// check ownership of nft
	owner := k.nftKeeper.GetOwner(ctx, nftId.ClassId, nftId.NftId)
	if owner.String() != sender.String() {
		return types.ErrNotNftOwner
	}
	params := k.GetParamSet(ctx)
	for !Contains(params.BidTokens, bidToken) {
		return types.ErrNotSupportedBidToken
	}

	return nil
}

func GetMemo(txBytes []byte, txCfg client.TxConfig) string {
	/// NOTE: this way requires txConfig by importing it into keeper struct
	txData, err := txCfg.TxDecoder()(txBytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)

		txData, err = txCfg.TxJSONDecoder()(txBytes)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}

	txBldr, err := txCfg.WrapTxBuilder(txData)
	if err != nil {
		return ""
	}
	memo := txBldr.GetTx().GetMemo()

	return memo
}
