package lockup

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/x/lockup/keeper"
	"github.com/osmosis-labs/osmosis/x/lockup/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetLastLockID(ctx, genState.LastLockId)
	if err := k.ResetAllLocks(ctx, genState.Locks); err != nil {
		return
	}
	if err := k.ResetAllSyntheticLocks(ctx, genState.SyntheticLocks); err != nil {
		return
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	locks, err := k.GetPeriodLocks(ctx)
	if err != nil {
		panic(err)
	}
	return &types.GenesisState{
		LastLockId:     k.GetLastLockID(ctx),
		Locks:          locks,
		SyntheticLocks: k.GetAllSyntheticLockups(ctx),
	}
}
