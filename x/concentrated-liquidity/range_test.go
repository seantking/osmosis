package concentrated_liquidity_test

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v16/app/apptesting"
	"github.com/osmosis-labs/osmosis/v16/x/concentrated-liquidity/math"
	"github.com/osmosis-labs/osmosis/v16/x/concentrated-liquidity/types"
)

func (s *KeeperTestSuite) TestMultipleRanges() {
	tests := map[string]struct {
		tickRanges      [][]int64
		rangeTestParams RangeTestParams
	}{
		"one range, default params": {
			tickRanges: [][]int64{
				{0, 10000},
			},
			rangeTestParams: DefaultRangeTestParams,
		},
		"one min width range": {
			tickRanges: [][]int64{
				{0, 100},
			},
			rangeTestParams: withTickSpacing(DefaultRangeTestParams, DefaultTickSpacing),
		},
		"two adjacent ranges": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: DefaultRangeTestParams,
		},
		"two adjacent ranges with current tick smaller than both": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, -20000),
		},
		"two adjacent ranges with current tick larger than both": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, 30000),
		},
		"two adjacent ranges with current tick exactly on lower bound": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, -10000),
		},
		"two adjacent ranges with current tick exactly between both": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, 10000),
		},
		"two adjacent ranges with current tick exactly on upper bound": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{10000, 20000},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, 20000),
		},
		"two non-adjacent ranges": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{20000, 30000},
			},
			rangeTestParams: DefaultRangeTestParams,
		},
		"two ranges with one tick gap in between, which is equal to current tick": {
			tickRanges: [][]int64{
				{799221, 799997},
				{799997 + 2, 812343},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, 799997+1),
		},
		"one range on large tick": {
			tickRanges: [][]int64{
				{207000000, 207000000 + 100},
			},
			rangeTestParams: withTickSpacing(DefaultRangeTestParams, DefaultTickSpacing),
		},
		"one position adjacent to left of current tick (no swaps)": {
			tickRanges: [][]int64{
				{-1, 0},
			},
			rangeTestParams: RangeTestParamsNoFuzzNoSwap,
		},
		"one position on left of current tick with gap (no swaps)": {
			tickRanges: [][]int64{
				{-2, -1},
			},
			rangeTestParams: RangeTestParamsNoFuzzNoSwap,
		},
		"one position adjacent to right of current tick (no swaps)": {
			tickRanges: [][]int64{
				{0, 1},
			},
			rangeTestParams: RangeTestParamsNoFuzzNoSwap,
		},
		"one position on right of current tick with gap (no swaps)": {
			tickRanges: [][]int64{
				{1, 2},
			},
			rangeTestParams: RangeTestParamsNoFuzzNoSwap,
		},
		"one range on small tick": {
			tickRanges: [][]int64{
				{-107000000, -107000000 + 100},
			},
			rangeTestParams: withDoubleFundedLP(DefaultRangeTestParams),
		},
		"one range on min tick": {
			tickRanges: [][]int64{
				{types.MinInitializedTick, types.MinInitializedTick + 100},
			},
			rangeTestParams: withDoubleFundedLP(DefaultRangeTestParams),
		},
		"initial current tick equal to min initialized tick": {
			tickRanges: [][]int64{
				{0, 1},
			},
			rangeTestParams: withCurrentTick(DefaultRangeTestParams, types.MinInitializedTick),
		},
		"three overlapping ranges with no swaps, current tick in one": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{0, 20000},
				{-7300, 12345},
			},
			rangeTestParams: withNoSwap(withCurrentTick(DefaultRangeTestParams, -9000)),
		},
		"three overlapping ranges with no swaps, current tick in two of three": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{0, 20000},
				{-7300, 12345},
			},
			rangeTestParams: withNoSwap(withCurrentTick(DefaultRangeTestParams, -7231)),
		},
		"three overlapping ranges with no swaps, current tick in all three": {
			tickRanges: [][]int64{
				{-10000, 10000},
				{0, 20000},
				{-7300, 12345},
			},
			rangeTestParams: withNoSwap(withCurrentTick(DefaultRangeTestParams, 109)),
		},
		/* TODO: uncomment when infinite loop bug is fixed
		"one range on max tick": {
			tickRanges: [][]int64{
				{types.MaxTick - 100, types.MaxTick},
			},
			rangeTestParams: withTickSpacing(DefaultRangeTestParams, DefaultTickSpacing),
		},
		"initial current tick equal to max tick": {
			tickRanges: [][]int64{
				{0, 1},
			},
			rangeTestParams: withCurrentTick(withTickSpacing(DefaultRangeTestParams, uint64(1)), types.MaxTick),
		},
		*/
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			s.runMultiplePositionRanges(tc.tickRanges, tc.rangeTestParams)
		})
	}
}

// runMultiplePositionRanges runs various test constructions and invariants on the given position ranges.
func (s *KeeperTestSuite) runMultiplePositionRanges(ranges [][]int64, rangeTestParams RangeTestParams) {
	// Preset seed to ensure deterministic test runs.
	rand.Seed(2)

	// TODO: add pool-related fuzz params (spread factor & number of pools)
	pool := s.PrepareCustomConcentratedPool(s.TestAccs[0], ETH, USDC, rangeTestParams.tickSpacing, rangeTestParams.spreadFactor)

	// Run full state determined by params while asserting invariants at each intermediate step
	s.setupRangesAndAssertInvariants(pool, ranges, rangeTestParams)

	// Assert global invariants on final state
	s.assertGlobalInvariants(ExpectedGlobalRewardValues{})
}

type RangeTestParams struct {
	// -- Base amounts --

	// Base number of assets for each position
	baseAssets sdk.Coins
	// Base number of positions for each range
	baseNumPositions int
	// Base amount to swap for each swap
	baseSwapAmount sdk.Int
	// Base amount to add after each new position
	baseTimeBetweenJoins time.Duration
	// Base incentive amount to have on each incentive record
	baseIncentiveAmount sdk.Int
	// Base emission rate per second for incentive
	baseEmissionRate sdk.Dec
	// Base denom for each incentive record (ID appended to this)
	baseIncentiveDenom string
	// List of addresses to swap from (randomly selected for each swap)
	numSwapAddresses int

	// -- Pool params --

	spreadFactor sdk.Dec
	tickSpacing  uint64

	// -- Fuzz params --

	fuzzAssets           bool
	fuzzNumPositions     bool
	fuzzSwapAmounts      bool
	fuzzTimeBetweenJoins bool
	fuzzIncentiveRecords bool

	// -- Optional additional test dimensions --

	// Have a single address for all positions in each range
	singleAddrPerRange bool
	// Create new active incentive records between each join
	newActiveIncentivesBetweenJoins bool
	// Create new inactive incentive records between each join
	newInactiveIncentivesBetweenJoins bool
	// Fund each position address with double the expected amount of assets.
	// Should only be used for cases where join amount gets pushed up due to
	// precision near min tick.
	doubleFundPositionAddr bool
	// Adjust input amounts for first position to set the starting current tick
	// to the given value.
	startingCurrentTick int64
}

func (r RangeTestParams) makeAddresses(totalPositions int, rangeLen int) []sdk.AccAddress {
	if r.singleAddrPerRange {
		return apptesting.CreateRandomAccounts(rangeLen)
	}
	return apptesting.CreateRandomAccounts(totalPositions)
}

var (
	DefaultRangeTestParams = RangeTestParams{
		// Base amounts
		baseNumPositions:     10,
		baseAssets:           sdk.NewCoins(sdk.NewCoin(ETH, sdk.NewInt(5000000000)), sdk.NewCoin(USDC, sdk.NewInt(5000000000))),
		baseTimeBetweenJoins: time.Hour,
		baseSwapAmount:       sdk.NewInt(10000000),
		numSwapAddresses:     10,
		baseIncentiveAmount:  sdk.NewInt(1000000000000000000),
		baseEmissionRate:     sdk.NewDec(1),
		baseIncentiveDenom:   "incentiveDenom",

		// Pool params
		spreadFactor: DefaultSpreadFactor,
		tickSpacing:  uint64(1),

		// Fuzz params
		fuzzNumPositions:     true,
		fuzzAssets:           true,
		fuzzSwapAmounts:      true,
		fuzzTimeBetweenJoins: true,
	}
	RangeTestParamsLargeSwap = RangeTestParams{
		// Base amounts
		baseNumPositions:     10,
		baseAssets:           sdk.NewCoins(sdk.NewCoin(ETH, sdk.NewInt(5000000000)), sdk.NewCoin(USDC, sdk.NewInt(5000000000))),
		baseTimeBetweenJoins: time.Hour,
		baseSwapAmount:       sdk.Int(sdk.MustNewDecFromStr("100000000000000000000000000000000000000")),
		numSwapAddresses:     10,
		baseIncentiveAmount:  sdk.NewInt(1000000000000000000),
		baseEmissionRate:     sdk.NewDec(1),
		baseIncentiveDenom:   "incentiveDenom",

		// Pool params
		spreadFactor: DefaultSpreadFactor,
		tickSpacing:  uint64(100),

		// Fuzz params
		fuzzNumPositions:     true,
		fuzzAssets:           true,
		fuzzTimeBetweenJoins: true,
	}
	RangeTestParamsNoFuzzNoSwap = RangeTestParams{
		// Base amounts
		baseNumPositions:     1,
		baseAssets:           sdk.NewCoins(sdk.NewCoin(ETH, sdk.NewInt(5000000000)), sdk.NewCoin(USDC, sdk.NewInt(5000000000))),
		baseTimeBetweenJoins: time.Hour,
		baseIncentiveAmount:  sdk.NewInt(1000000000000000000),
		baseEmissionRate:     sdk.NewDec(1),
		baseIncentiveDenom:   "incentiveDenom",

		// Pool params
		spreadFactor: DefaultSpreadFactor,
		tickSpacing:  uint64(1),
	}
)

func withDoubleFundedLP(params RangeTestParams) RangeTestParams {
	params.doubleFundPositionAddr = true
	return params
}

func withCurrentTick(params RangeTestParams, tick int64) RangeTestParams {
	params.startingCurrentTick = tick
	return params
}

func withTickSpacing(params RangeTestParams, tickSpacing uint64) RangeTestParams {
	params.tickSpacing = tickSpacing
	return params
}

func withNoSwap(params RangeTestParams) RangeTestParams {
	params.baseSwapAmount = sdk.Int{}
	return params
}

func (s *KeeperTestSuite) setupRanges(pool types.ConcentratedPoolExtension, ranges [][]int64, testParams RangeTestParams) (int, []int, []sdk.AccAddress, []sdk.AccAddress) {
	// Prepare a slice tracking how many positions to create on each range.
	// setup addresses as well.
	numPositionSlice, totalPositions := s.prepareNumPositionSlice(ranges, testParams.baseNumPositions, testParams.fuzzNumPositions)
	positionAddresses := testParams.makeAddresses(totalPositions, len(ranges))
	swapAddresses := apptesting.CreateRandomAccounts(testParams.numSwapAddresses)

	// --- Incentive setup ---

	if testParams.baseIncentiveAmount != (sdk.Int{}) {
		incentiveAddr := apptesting.CreateRandomAccounts(1)[0]
		incentiveAmt := testParams.baseIncentiveAmount
		emissionRate := testParams.baseEmissionRate
		incentiveCoin := sdk.NewCoin(fmt.Sprintf("%s%d", testParams.baseIncentiveDenom, 0), incentiveAmt)
		s.FundAcc(incentiveAddr, sdk.NewCoins(incentiveCoin))
		_, err := s.clk.CreateIncentive(s.Ctx, pool.GetId(), incentiveAddr, incentiveCoin, emissionRate, s.Ctx.BlockTime(), types.DefaultAuthorizedUptimes[0])
		s.Require().NoError(err)
	}

	return totalPositions, numPositionSlice, positionAddresses, swapAddresses
}

// setupRangesAndAssertInvariants sets up the state specified by `testParams` on the given set of ranges.
// It also asserts global invariants at each intermediate step.
func (s *KeeperTestSuite) setupRangesAndAssertInvariants(pool types.ConcentratedPoolExtension, ranges [][]int64, testParams RangeTestParams) {
	totalPositions, numPositionSlice, positionAddresses, swapAddresses := s.setupRanges(pool, ranges, testParams)

	// --- Position setup ---

	// This loop runs through each given tick range and does the following at each iteration:
	// 1. Set up a position
	// 2. Let time elapse
	// 3. Execute a swap
	totalLiquidity, totalAssets, totalTimeElapsed, allPositionIds, lastVisitedBlockIndex, cumulativeEmittedIncentives, lastIncentiveTrackerUpdate := sdk.ZeroDec(), sdk.NewCoins(), time.Duration(0), []uint64{}, 0, sdk.DecCoins{}, s.Ctx.BlockTime()
	for curRange := range ranges {
		curBlock := 0
		startNumPositions := len(allPositionIds)
		for curNumPositions := lastVisitedBlockIndex; curNumPositions < lastVisitedBlockIndex+numPositionSlice[curRange]; curNumPositions++ {
			// By default we create a new address for each position, but if the test params specify using a single address
			// for each range, we handle that logic here.
			var curAddr sdk.AccAddress
			if testParams.singleAddrPerRange {
				// If we are using a single address per range, we use the address corresponding to the current range.
				curAddr = positionAddresses[curRange]
			} else {
				// If we're not using a single address per range, we use a unique address for each position.
				curAddr = positionAddresses[curNumPositions]
			}

			// Set up assets for new position
			curAssets := getRandomizedAssets(testParams.baseAssets, testParams.fuzzAssets)

			// If a desired current tick was specified, retrieve special asset amounts for the first position
			if testParams.startingCurrentTick != 0 && curNumPositions == 0 {
				curAssets = s.getInitialPositionAssets(pool, testParams.startingCurrentTick)
			}

			roundingError := sdk.NewCoins(sdk.NewCoin(pool.GetToken0(), sdk.OneInt()), sdk.NewCoin(pool.GetToken1(), sdk.OneInt()))
			s.FundAcc(curAddr, curAssets.Add(roundingError...))

			// Double fund LP address if applicable
			if testParams.doubleFundPositionAddr {
				s.FundAcc(curAddr, curAssets.Add(roundingError...))
			}

			// TODO: implement intermediate record creation with fuzzing

			// Track emitted incentives here
			cumulativeEmittedIncentives, lastIncentiveTrackerUpdate = s.trackEmittedIncentives(cumulativeEmittedIncentives, lastIncentiveTrackerUpdate)

			// Set up position
			curPositionId, actualAmt0, actualAmt1, curLiquidity, actualLowerTick, actualUpperTick, err := s.clk.CreatePosition(s.Ctx, pool.GetId(), curAddr, curAssets, sdk.ZeroInt(), sdk.ZeroInt(), ranges[curRange][0], ranges[curRange][1])
			s.Require().NoError(err)

			// Ensure position was set up correctly and didn't break global invariants
			s.Require().Equal(ranges[curRange][0], actualLowerTick)
			s.Require().Equal(ranges[curRange][1], actualUpperTick)
			s.assertGlobalInvariants(ExpectedGlobalRewardValues{})

			// Let time elapse after join if applicable
			timeElapsed := s.addRandomizedBlockTime(testParams.baseTimeBetweenJoins, testParams.fuzzTimeBetweenJoins)

			// Execute swap against pool if applicable
			fmt.Println("-------------------- Begin new Swap --------------------")
			cctx, write := s.Ctx.CacheContext()
			swappedIn, swappedOut, ok := s.executeRandomizedSwap(cctx, pool, swapAddresses, testParams.baseSwapAmount, testParams.fuzzSwapAmounts)
			if !ok {
				continue
			}
			write()
			s.assertGlobalInvariants(ExpectedGlobalRewardValues{})

			// Track changes to state
			actualAddedCoins := sdk.NewCoins(sdk.NewCoin(pool.GetToken0(), actualAmt0), sdk.NewCoin(pool.GetToken1(), actualAmt1))
			totalAssets = totalAssets.Add(actualAddedCoins...)
			if testParams.baseSwapAmount != (sdk.Int{}) && (swappedIn != sdk.Coin{} || swappedOut != sdk.Coin{}) {
				totalAssets = totalAssets.Add(swappedIn).Sub(sdk.NewCoins(swappedOut))
			}
			totalLiquidity = totalLiquidity.Add(curLiquidity)
			totalTimeElapsed = totalTimeElapsed + timeElapsed
			allPositionIds = append(allPositionIds, curPositionId)
			curBlock++
		}
		endNumPositions := len(allPositionIds)

		// Ensure the correct number of positions were set up in current range
		s.Require().Equal(numPositionSlice[curRange], endNumPositions-startNumPositions, "Incorrect number of positions set up in range %d", curRange)

		lastVisitedBlockIndex += curBlock
	}

	// Ensure that the correct number of positions were set up globally
	s.Require().Equal(totalPositions, len(allPositionIds))

	// Ensure the pool balance is exactly equal to the assets added + amount swapped in - amount swapped out
	poolAssets := s.App.BankKeeper.GetAllBalances(s.Ctx, pool.GetAddress())
	poolSpreadRewards := s.App.BankKeeper.GetAllBalances(s.Ctx, pool.GetSpreadRewardsAddress())
	// We rebuild coins to handle nil cases cleanly
	s.Require().Equal(sdk.NewCoins(totalAssets...), sdk.NewCoins(poolAssets.Add(poolSpreadRewards...)...))

	// Do a final checkpoint for incentives and then run assertions on expected global claimable value
	cumulativeEmittedIncentives, lastIncentiveTrackerUpdate = s.trackEmittedIncentives(cumulativeEmittedIncentives, lastIncentiveTrackerUpdate)
	truncatedEmissions, _ := cumulativeEmittedIncentives.TruncateDecimal()

	// Run global assertions with an optional parameter specifying the expected incentive amount claimable by all positions.
	// We specifically need to do this for incentives because all the emissions are pre-loaded into the incentive address, making
	// balance assertions pass trivially in most cases.
	s.assertGlobalInvariants(ExpectedGlobalRewardValues{TotalIncentives: truncatedEmissions})
}

// numPositionSlice prepares a slice tracking the number of positions to create on each range, fuzzing the number at each step if applicable.
// Returns a slice representing the number of positions for each range index.
//
// We run this logic in a separate function for two main reasons:
// 1. Simplify position setup logic by fuzzing the number of positions upfront, letting us loop through the positions to set them up
// 2. Abstract as much fuzz logic from the core setup loop, which is already complex enough as is
func (s *KeeperTestSuite) prepareNumPositionSlice(ranges [][]int64, baseNumPositions int, fuzzNumPositions bool) ([]int, int) {
	// Create slice representing number of positions for each range index.
	// Default case is `numPositions` on each range unless fuzzing is turned on.
	numPositionsPerRange := make([]int, len(ranges))
	totalPositions := 0

	// Loop through each range and set number of positions, fuzzing if applicable.
	for i := range ranges {
		numPositionsPerRange[i] = baseNumPositions

		// If applicable, fuzz the number of positions on current range
		if fuzzNumPositions {
			// Fuzzed amount should be between 1 and (2 * numPositions) + 1 (up to 100% fuzz both ways from numPositions)
			numPositionsPerRange[i] = int(fuzzInt64(int64(baseNumPositions), 2))
		}

		// Track total positions
		totalPositions += numPositionsPerRange[i]
	}

	return numPositionsPerRange, totalPositions
}

// executeRandomizedSwap executes a swap against the pool, fuzzing the swap amount if applicable.
// The direction of the swap is chosen randomly, but the swap function used is always SwapInGivenOut to
// ensure it is always possible to swap against the pool without having to use lower level calc functions.
// TODO: Make swaps that target getting to a tick boundary exactly
func (s *KeeperTestSuite) executeRandomizedSwap(ctx sdk.Context, pool types.ConcentratedPoolExtension, swapAddresses []sdk.AccAddress, baseSwapAmount sdk.Int, fuzzSwap bool) (sdk.Coin, sdk.Coin, bool) {
	// Quietly skip if no swap assets or swap addresses provided
	if (baseSwapAmount == sdk.Int{}) || len(swapAddresses) == 0 {
		return sdk.Coin{}, sdk.Coin{}, false
	}

	poolLiquidity := s.App.BankKeeper.GetAllBalances(ctx, pool.GetAddress())
	s.Require().True(len(poolLiquidity) == 1 || len(poolLiquidity) == 2, "Pool liquidity should be in one or two tokens")

	// Choose swap address
	swapAddressIndex := fuzzInt64(int64(len(swapAddresses)-1), 1)
	swapAddress := swapAddresses[swapAddressIndex]

	// Decide which denom to swap in & out

	// var swapInDenom, swapOutDenom string
	// if len(poolLiquidity) == 1 {
	// 	// If all pool liquidity is in one token, swap in the other token
	// 	swapOutDenom = poolLiquidity[0].Denom
	// 	if swapOutDenom == pool.GetToken0() {
	// 		swapInDenom = pool.GetToken1()
	// 	} else {
	// 		swapInDenom = pool.GetToken0()
	// 	}
	// } else {
	// 	// Otherwise, randomly determine which denom to swap in & out
	// 	swapInDenom, swapOutDenom = randOrder(pool.GetToken0(), pool.GetToken1())
	// }

	updatedPool, err := s.clk.GetPoolById(ctx, pool.GetId())
	s.Require().NoError(err)
	// TODO: allow target tick to be specified and fuzzed

	// Note: the early return here was simply to rush repro the panic. This logic will ultimately live in separate branches depending on whether
	// testParams.swapToTickBoundary is enabled or not.
	return s.executeSwapToTickBoundary(ctx, updatedPool, swapAddress, updatedPool.GetCurrentTick()+1, false)

	// // TODO: pick a more granular amount to fund without losing ability to swap at really high/low ticks
	// swapInFunded := sdk.NewCoin(swapInDenom, sdk.Int(sdk.MustNewDecFromStr("10000000000000000000000000000000000000000")))
	// s.FundAcc(swapAddress, sdk.NewCoins(swapInFunded))

	// baseSwapOutAmount := sdk.MinInt(baseSwapAmount, poolLiquidity.AmountOf(swapOutDenom).ToDec().Mul(sdk.MustNewDecFromStr("0.5")).TruncateInt())
	// if fuzzSwap {
	// 	// Fuzz +/- 100% of base swap amount
	// 	baseSwapOutAmount = sdk.NewInt(fuzzInt64(baseSwapOutAmount.Int64(), 2))
	// }

	// swapOutCoin := sdk.NewCoin(swapOutDenom, baseSwapOutAmount)

	// // If the swap we're about to execute will not generate enough input, we skip the swap.
	// if swapOutDenom == pool.GetToken1() {
	// 	pool, err := s.clk.GetPoolById(s.Ctx, pool.GetId())
	// 	s.Require().NoError(err)

	// 	poolSpotPrice := pool.GetCurrentSqrtPrice().Power(osmomath.NewBigDec(2))
	// 	minSwapOutAmount := poolSpotPrice.Mul(osmomath.SmallestDec()).SDKDec().TruncateInt()
	// 	poolBalances := s.App.BankKeeper.GetAllBalances(s.Ctx, pool.GetAddress())
	// 	if poolBalances.AmountOf(swapOutDenom).LTE(minSwapOutAmount) {
	// 		return sdk.Coin{}, sdk.Coin{}
	// 	}
	// }

	// // Note that we set the price limit to zero to ensure that the swap can execute in either direction (gets automatically set to correct limit)
	// swappedIn, swappedOut, _, err = s.clk.SwapInAmtGivenOut(s.Ctx, swapAddress, pool, swapOutCoin, swapInDenom, pool.GetSpreadFactor(s.Ctx), sdk.ZeroDec())
	// s.Require().NoError(err)

	// return swappedIn, swappedOut
}

// executeSwapToTickBoundary executes a swap against the pool to get to the specified tick boundary, randomizing the chosen tick if applicable.
func (s *KeeperTestSuite) executeSwapToTickBoundary(ctx sdk.Context, pool types.ConcentratedPoolExtension, swapAddress sdk.AccAddress, targetTick int64, fuzzTick bool) (sdk.Coin, sdk.Coin, bool) {
	// zeroForOne := swapInDenom == pool.GetToken0()

	pool, err := s.clk.GetPoolById(s.Ctx, pool.GetId())
	s.Require().NoError(err)
	fmt.Println("current tick: ", pool.GetCurrentTick())
	currentTick := pool.GetCurrentTick()
	zeroForOne := currentTick >= targetTick
	amountInRequired, _, _ := s.computeSwapAmounts(pool.GetId(), pool.GetCurrentSqrtPrice(), targetTick, zeroForOne, false)

	var swapInDenom, swapOutDenom string
	if zeroForOne {
		swapInDenom = pool.GetToken0()
		swapOutDenom = pool.GetToken1()
	} else {
		swapInDenom = pool.GetToken1()
		swapOutDenom = pool.GetToken0()
	}

	poolSpotPrice := pool.GetCurrentSqrtPrice().Power(osmomath.NewBigDec(2))
	minSwapOutAmount := poolSpotPrice.Mul(osmomath.SmallestDec()).SDKDec().TruncateInt()
	poolBalances := s.App.BankKeeper.GetAllBalances(ctx, pool.GetAddress())
	if poolBalances.AmountOf(swapOutDenom).LTE(minSwapOutAmount) {
		fmt.Println("skipped")
		return sdk.Coin{}, sdk.Coin{}, false
	}

	fmt.Println("dec amt in required to get to tick boundary: ", amountInRequired)
	swapInFunded := sdk.NewCoin(swapInDenom, amountInRequired.TruncateInt())
	s.FundAcc(swapAddress, sdk.NewCoins(swapInFunded))

	// Execute swap
	fmt.Println("begin keeper swap")
	swappedIn, swappedOut, _, err := s.clk.SwapOutAmtGivenIn(ctx, swapAddress, pool, swapInFunded, swapOutDenom, pool.GetSpreadFactor(s.Ctx), sdk.ZeroDec())
	if errors.As(err, &types.InvalidAmountCalculatedError{}) {
		// If the swap we're about to execute will not generate enough output, we skip the swap.
		// it would error for a real user though. This is good though, since that user would just be burning funds.
		if err.(types.InvalidAmountCalculatedError).Amount.IsZero() {
			return sdk.Coin{}, sdk.Coin{}, false
		} else {
			s.Require().NoError(err)
		}
	} else {
		s.Require().NoError(err)
	}

	return swappedIn, swappedOut, true
}

func randOrder[T any](a, b T) (T, T) {
	if rand.Int()%2 == 0 {
		return a, b
	}
	return b, a
}

// addRandomizedBlockTime adds the given block time to the context, fuzzing the added time if applicable.
func (s *KeeperTestSuite) addRandomizedBlockTime(baseTimeToAdd time.Duration, fuzzTime bool) time.Duration {
	if baseTimeToAdd != time.Duration(0) {
		timeToAdd := baseTimeToAdd
		if fuzzTime {
			// Fuzz +/- 100% of base time to add
			timeToAdd = time.Duration(fuzzInt64(int64(baseTimeToAdd), 2))
		}

		s.AddBlockTime(timeToAdd)
	}

	return baseTimeToAdd
}

// trackEmittedIncentives takes in a cumulative incentives distributed and the last time this number was updated.
// CONTRACT: cumulativeTrackedIncentives has been updated immediately before each new incentive record that was created
func (s *KeeperTestSuite) trackEmittedIncentives(cumulativeTrackedIncentives sdk.DecCoins, lastTrackerUpdateTime time.Time) (sdk.DecCoins, time.Time) {
	// Fetch all incentive records across all pools
	allPools, err := s.clk.GetPools(s.Ctx)
	s.Require().NoError(err)
	allIncentiveRecords := make([]types.IncentiveRecord, 0)
	for _, pool := range allPools {
		curPoolRecords, err := s.clk.GetAllIncentiveRecordsForPool(s.Ctx, pool.GetId())
		s.Require().NoError(err)

		allIncentiveRecords = append(allIncentiveRecords, curPoolRecords...)
	}

	// Track new emissions since last checkpoint, factoring in when each incentive record started emitting
	updatedTrackedIncentives := cumulativeTrackedIncentives
	for _, incentiveRecord := range allIncentiveRecords {
		recordStartTime := incentiveRecord.IncentiveRecordBody.StartTime

		// If the record hasn't started emitting yet, skip it
		if recordStartTime.After(s.Ctx.BlockTime()) {
			continue
		}

		secondsEmitted := sdk.ZeroDec()
		if recordStartTime.Before(lastTrackerUpdateTime) {
			// If the record started emitting prior to the last incentiveCreationTime (the last time we checkpointed),
			// then we assume it has been emitting for the whole period since then.
			secondsEmitted = sdk.NewDec(int64(s.Ctx.BlockTime().Sub(lastTrackerUpdateTime))).QuoInt64(int64(time.Second))
		} else if recordStartTime.Before(s.Ctx.BlockTime()) {
			// If the record started emitting between the last incentiveCreationTime and now, then we only track the
			// emissions between when it started and now.
			secondsEmitted = sdk.NewDec(int64(s.Ctx.BlockTime().Sub(recordStartTime))).QuoInt64(int64(time.Second))
		}

		emissionRate := incentiveRecord.IncentiveRecordBody.EmissionRate
		incentiveDenom := incentiveRecord.IncentiveRecordBody.RemainingCoin.Denom

		// Track emissions for the current record
		emittedAmount := emissionRate.Mul(secondsEmitted)
		emittedDecCoin := sdk.NewDecCoinFromDec(incentiveDenom, emittedAmount)
		updatedTrackedIncentives = updatedTrackedIncentives.Add(emittedDecCoin)
	}

	return updatedTrackedIncentives, s.Ctx.BlockTime()
}

// getInitialPositionAssets returns the assets required for the first position in a pool to set the initial current tick to the given value.
func (s *KeeperTestSuite) getInitialPositionAssets(pool types.ConcentratedPoolExtension, initialCurrentTick int64) sdk.Coins {
	requiredPrice, err := math.TickToPrice(initialCurrentTick)
	s.Require().NoError(err)

	// Calculate asset amounts that would be required to get the required spot price (rounding up on asset1 to ensure we stay in the intended tick)
	asset0Amount := sdk.NewInt(100000000000000)
	asset1Amount := sdk.NewDecFromInt(asset0Amount).Mul(requiredPrice).Ceil().TruncateInt()

	assetCoins := sdk.NewCoins(
		sdk.NewCoin(pool.GetToken0(), asset0Amount),
		sdk.NewCoin(pool.GetToken1(), asset1Amount),
	)

	return assetCoins
}

// getFuzzedAssets returns the base asset amount, fuzzing each asset if applicable
func getRandomizedAssets(baseAssets sdk.Coins, fuzzAssets bool) sdk.Coins {
	finalAssets := baseAssets
	if fuzzAssets {
		fuzzedAssets := make([]sdk.Coin, len(baseAssets))
		for coinIndex, coin := range baseAssets {
			// Fuzz +/- 100% of current amount
			newAmount := fuzzInt64(coin.Amount.Int64(), 2)
			fuzzedAssets[coinIndex] = sdk.NewCoin(coin.Denom, sdk.NewInt(newAmount))
		}

		finalAssets = fuzzedAssets
	}

	return finalAssets
}

// fuzzInt64 fuzzes an int64 number uniformly within a range defined by `multiplier` and centered on the provided `intToFuzz`.
func fuzzInt64(intToFuzz int64, multiplier int64) int64 {
	return (rand.Int63() % (multiplier * intToFuzz)) + 1
}
