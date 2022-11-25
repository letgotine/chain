package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	nftkeeper "github.com/cosmos/cosmos-sdk/x/nft/keeper"

	simapp "github.com/UnUniFi/chain/app"
	appparams "github.com/UnUniFi/chain/app/params"
	"github.com/UnUniFi/chain/x/nftmarket/keeper"
	"github.com/UnUniFi/chain/x/nftmarket/types"
)

var (
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		distrtypes.ModuleName:      nil,
		minttypes.ModuleName:       {authtypes.Minter},
		// liquiditytypes.ModuleName:      {authtypes.Minter, authtypes.Burner},
		// ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		// wasm.ModuleName:             {authtypes.Burner},
		nft.ModuleName:      nil,
		types.ModuleName:    nil,
		types.NftTradingFee: nil,
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	app         *simapp.App
	addrs       []sdk.AccAddress
	queryClient types.QueryClient
	// keeper      keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	app := simapp.Setup(suite.T(), isCheckTx)

	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	suite.addrs = simapp.AddTestAddrsIncremental(app, suite.ctx, 3, sdk.NewInt(30000000))
	suite.app = app
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.NftmarketKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)

	encodingConfig := appparams.MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler

	txCfg := encodingConfig.TxConfig
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, banktypes.StoreKey, types.StoreKey, types.MemStoreKey, nft.StoreKey)
	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix,
	)
	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.BlockedAddrs(),
	)
	nftKeeper := nftkeeper.NewKeeper(keys[nft.StoreKey], appCodec, accountKeeper, bankKeeper)
	keeper := keeper.NewKeeper(appCodec, txCfg, keys[types.StoreKey], keys[types.MemStoreKey], suite.app.GetSubspace(types.ModuleName), accountKeeper, bankKeeper, nftKeeper)
	hooks := dummyNftmarketHook{}
	keeper.SetHooks(&hooks)
	// suite.keeper = keeper
	// suite.app.NftmarketKeeper = keeper
}

func (suite *KeeperTestSuite) TestSuiteKeeper() {
	suite.SetupTest()
}
func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
