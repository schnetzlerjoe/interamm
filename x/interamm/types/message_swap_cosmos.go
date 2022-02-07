package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwapCosmos = "swap_cosmos"

var _ sdk.Msg = &MsgSwapCosmos{}

func NewMsgSwapCosmos(creator string, poolId string, swapType string, offerCoin string, demandCoinDenom string, orderPrice string, swapFeeRate string) *MsgSwapCosmos {
	return &MsgSwapCosmos{
		Creator:         creator,
		PoolId:          poolId,
		SwapType:        swapType,
		OfferCoin:       offerCoin,
		DemandCoinDenom: demandCoinDenom,
		OrderPrice:      orderPrice,
		SwapFeeRate:     swapFeeRate,
	}
}

func (msg *MsgSwapCosmos) Route() string {
	return RouterKey
}

func (msg *MsgSwapCosmos) Type() string {
	return TypeMsgSwapCosmos
}

func (msg *MsgSwapCosmos) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapCosmos) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwapCosmos) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
