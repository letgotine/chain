package keeper

import (
	"fmt"
	"strconv"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/interchainquery/types"
)

func GenerateQueryHash(connection_id string, chain_id string, query_type string, request []byte, module string, height int64) string {
	return fmt.Sprintf("%x", crypto.Sha256(append([]byte(module+connection_id+chain_id+query_type+strconv.FormatInt(height, 10)), request...)))
}

// ----------------------------------------------------------------

func (k Keeper) NewQuery(ctx sdk.Context, module string, connection_id string, chain_id string, query_type string, request []byte, period sdk.Int, callback_id string, ttl uint64, height int64) *types.Query {
	return &types.Query{
		Id:           GenerateQueryHash(connection_id, chain_id, query_type, request, module, height),
		ConnectionId: connection_id,
		ChainId:      chain_id,
		QueryType:    query_type,
		Request:      request,
		Period:       period,
		LastHeight:   sdk.ZeroInt(),
		CallbackId:   callback_id,
		Ttl:          ttl,
		Height:       height,
	}
}

// GetQuery returns query
func (k Keeper) GetQuery(ctx sdk.Context, id string) (types.Query, bool) {
	query := types.Query{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	bz := store.Get([]byte(id))
	if len(bz) == 0 {
		return query, false
	}
	k.cdc.MustUnmarshal(bz, &query)
	return query, true
}

// SetQuery set query info
func (k Keeper) SetQuery(ctx sdk.Context, query types.Query) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	bz := k.cdc.MustMarshal(&query)
	store.Set([]byte(query.Id), bz)
}

// DeleteQuery delete query info
func (k Keeper) DeleteQuery(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	store.Delete([]byte(id))
}

// IterateQueries iterate through queries
func (k Keeper) IterateQueries(ctx sdk.Context, fn func(index int64, queryInfo types.Query) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixQuery)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		query := types.Query{}
		k.cdc.MustUnmarshal(iterator.Value(), &query)
		stop := fn(i, query)

		if stop {
			break
		}
		i++
	}
}

// AllQueries returns every queryInfo in the store
func (k Keeper) AllQueries(ctx sdk.Context) []types.Query {
	queries := []types.Query{}
	k.IterateQueries(ctx, func(_ int64, queryInfo types.Query) (stop bool) {
		queries = append(queries, queryInfo)
		return false
	})
	return queries
}
