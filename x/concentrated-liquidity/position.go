package concentrated_liquidity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/osmoutils"
	types "github.com/osmosis-labs/osmosis/v12/x/concentrated-liquidity/types"
)

// func (k Keeper) UpdateTickWithNewLiquidity(ctx sdk.Context, poolId uint64, tickIndex sdk.Int, liquidityDelta sdk.Int) {
// 	tickInfo := k.getTickInfoByPoolIDAndTickIndex(ctx, poolId, tickIndex)

// 	liquidityBefore := tickInfo.Liquidity
// 	liquidityAfter := liquidityBefore.Add(liquidityDelta)

// 	tickInfo.Liquidity = liquidityAfter

// 	if liquidityBefore == sdk.ZeroInt() {
// 		tickInfo.Initialized = true
// 	}

// 	k.setTickInfoByPoolIDAndTickIndex(ctx, poolId, tickIndex, tickInfo)
// }

func (k Keeper) UpdatePositionWithLiquidity(ctx sdk.Context,
	poolId uint64,
	owner string,
	lowerTick, upperTick sdk.Int,
	liquidityDelta sdk.Int) {
	position := k.getPosition(ctx, poolId, owner, lowerTick, upperTick)

	liquidityBefore := position.Liquidity
	liquidityAfter := liquidityBefore.Add(liquidityDelta)
	position.Liquidity = liquidityAfter

	k.setPosition(ctx, poolId, owner, lowerTick, upperTick, position)
}

func (k Keeper) getPosition(ctx sdk.Context, poolId uint64, owner string, lowerTick, upperTick sdk.Int) Position {
	store := ctx.KVStore(k.storeKey)
	position := Position{}
	key := types.KeyPosition(poolId, owner, lowerTick, upperTick)
	osmoutils.MustGet(store, key, &position)
	return position
}

func (k Keeper) setPosition(ctx sdk.Context,
	poolId uint64,
	owner string,
	lowerTick, upperTick sdk.Int,
	position Position) {

	store := ctx.KVStore(k.storeKey)

	key := types.KeyPosition(poolId, owner, lowerTick, upperTick)
	osmoutils.MustSet(store, key, &position)
}
