package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSwapOsmosis = "swap_osmosis"

var _ sdk.Msg = &MsgSwapOsmosis{}

func NewMsgSwapOsmosis(creator string, tokenIn string, tokenOutMinAmount string) *MsgSwapOsmosis {
	return &MsgSwapOsmosis{
		Creator:           creator,
		TokenIn:           tokenIn,
		TokenOutMinAmount: tokenOutMinAmount,
	}
}

func (msg *MsgSwapOsmosis) Route() string {
	return RouterKey
}

func (msg *MsgSwapOsmosis) Type() string {
	return TypeMsgSwapOsmosis
}

func (msg *MsgSwapOsmosis) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSwapOsmosis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSwapOsmosis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
