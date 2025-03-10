//go:build !cgo

package keeper

import (
	corestoretypes "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

// NewKeeper creates a new contract Keeper instance
// If customEncoders is non-nil, we can use this to override some of the message handler, especially custom
func NewKeeper(
	cdc codec.Codec,
	storeService corestoretypes.KVStoreService,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distrKeeper types.DistributionKeeper,
	ics4Wrapper types.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portSource types.ICS20TransferPortSource,
	router MessageRouter,
	_ GRPCQueryRouter,
	homeDir string,
	nodeConfig types.NodeConfig,
	vmConfig types.VMConfig,
	availableCapabilities []string,
	authority string,
	opts ...Option,
) Keeper {
	panic("not implemented, please build with cgo enabled")
}
