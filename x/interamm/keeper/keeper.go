package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/schnetzlerjoe/interamm/x/interamm/types"
	liquiditytypes "github.com/tendermint/liquidity/x/liquidity/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		icaControllerKeeper icacontrollerkeeper.Keeper
		scopedKeeper        capabilitykeeper.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Keeper that creates a Cosmos Hub send message from inputs and submits it to Cosmos Hub as ICA Tx
func (k Keeper) ICASend(ctx sdk.Context, sourcePort string, sourceChannel string, token sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, owner string, connectionId string) error {

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionId, portID)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	sendmsg := transfertypes.NewMsgTransfer(sourcePort, sourceChannel, token, sender, receiver, timeoutHeight, timeoutTimestamp)

	data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{sendmsg})
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp = uint64(time.Now().Add(time.Minute).UnixNano())
	_, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connectionId, portID, packetData, timeoutTimestamp)
	if err != nil {
		return err
	}

	return nil
}

// Keeper that creates a Cosmos Hub IBC transfer message from inputs and submits it to Cosmos Hub as ICA Tx
func (k Keeper) ICATransfer(ctx sdk.Context, sourcePort string, sourceChannel string, token sdk.Coin, sender string, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64, owner string, connectionId string) error {

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionId, portID)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	sendmsg := transfertypes.NewMsgTransfer(sourcePort, sourceChannel, token, sender, receiver, timeoutHeight, timeoutTimestamp)

	data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{sendmsg})
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp = uint64(time.Now().Add(time.Minute).UnixNano())
	_, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connectionId, portID, packetData, timeoutTimestamp)
	if err != nil {
		return err
	}

	return nil
}

// Keeper that creates a GDex Cosmos swap message from inputs and submits it to Cosmos Hub as ICA Tx
func (k Keeper) CosmosSwap(ctx sdk.Context, swapRequester sdk.AccAddress, poolID uint64, swapTypeID uint32, offerCoin sdk.Coin, demandCoinDenom string, orderPrice sdk.Dec, swapFeeRate sdk.Dec, owner string, connectionId string) error {

	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, connectionId, portID)
	if !found {
		return sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	swapmsg := liquiditytypes.NewMsgSwapWithinBatch(swapRequester, poolID, swapTypeID, offerCoin, demandCoinDenom, orderPrice, swapFeeRate)

	data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{swapmsg})
	if err != nil {
		return err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := time.Now().Add(time.Minute).UnixNano()
	_, err = k.icaControllerKeeper.SendTx(ctx, chanCap, connectionId, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	return nil
}
